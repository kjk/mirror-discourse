package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kjk/common/u"
)

// var banner_html = `<div style="color: red; padding-top: 1em; font-size: 20pt; font-weight: bold;">
// <center>
// Forum moved <a href="https://github.com/sumatrapdfreader/sumatrapdf/discussions">here</a>!
// </center>
// </div>`

var banner_html = ""

// This is a port of https://www.marksmath.org/ArchiveDiscourse/
// https://meta.discourse.org/t/a-basic-discourse-archival-tool/62614

// A tool to archive Discourse forum by using their .json data.
// I tried traditional mirroring with wget and HTTrack and
// they didn't work well

//go:embed missing_image.png
var missing_image_png []byte

//go:embed site-logo.png
var site_logo_png []byte

//go:embed main.css
var css []byte

// Template for the main page. Subsequent code will replace a few items indicated by
//go:embed tmpl_main.html
var main_template string

// Template for the individual topic pages
//go:embed tmpl_topic.html
var topic_template string

var base_url = ""
var base_scheme = ""
var archive_blurb = ""
var dstDir = "meta_discourse"
var imagesDir = filepath.Join(dstDir, "images")
var site_title = "Dummy title"

// Function that writes out each individual topic page
func write_topic(topic_json *Topic) {
	topic_download_url := base_url + "/t/" + topic_json.Slug + "/" + strconv.Itoa(topic_json.ID)
	topic_relative_dir := filepath.Join(dstDir, "t", topic_json.Slug, strconv.Itoa(topic_json.ID))
	err := os.MkdirAll(topic_relative_dir, 0755)
	must(err)
	var topic TopicResponse
	requests_get_json_must(topic_download_url+".json", &topic)
	posts_json := topic.PostStream.Posts
	post_list_string := ""
	for _, post_json := range posts_json {
		post_list_string = post_list_string + post_row(post_json)
	}
	topic_file_string := strings.ReplaceAll(topic_template, "<!-- TOPIC_TITLE -->", topic_json.FancyTitle)
	topic_file_string = strings.ReplaceAll(topic_file_string, "<!-- BANNER_HTML -->", banner_html)
	topic_file_string = strings.ReplaceAll(topic_file_string, "<!-- JUST_SITE_TITLE -->", site_title)

	topic_file_string = strings.ReplaceAll(topic_file_string, "<!-- ARCHIVE_BLURB -->", archive_blurb)
	topic_file_string = strings.ReplaceAll(topic_file_string, "<!-- POST_LIST -->", post_list_string)

	topic_path := filepath.Join(topic_relative_dir, "index.html")
	os.WriteFile(topic_path, []byte(topic_file_string), 0644)
}

func writeURLToFileMust(uri string, path string) {
	if u.FileExists(path) {
		logf("writeURLToFileMust: '%s' already exists\n", path)
		return
	}
	response := requests_get_must(uri)
	writeFileMust(path, response)
}

func writeURLToFile(uri string, path string) {
	if u.FileExists(path) {
		logf("writeURLToFile: '%s' already exists\n", path)
		return
	}
	response, err := requests_get(uri)
	if err != nil {
		return
	}
	writeFile(path, response)
}

func postBodyTransform(content string) string {
	r := bytes.NewBufferString(content)
	soup, err := goquery.NewDocumentFromReader(r)
	must(err)
	// Since we don't generate user information,
	// replace any anchors of class mention with a span
	// TODO: implement me
	// mention_tags := soup.Find("a.mention")
	// for _, tag := range mention_tags.Nodes {
	// 	//tag.Parent.RemoveChild()
	// 	/*
	// 	   try:
	// 	       rep = bs('<span class="mention"></span>', "html.parser").find('span')
	// 	       rep.string = tag.string
	// 	       tag.replaceWith(rep)
	// 	   except TypeError:
	// 	       pass
	// 	*/
	// }
	img_tags := soup.Find("img")
	img_tags.Text()
	for _, img_tag := range img_tags.Nodes {
		img_url := getAttrMust(img_tag.Attr, "src")
		parsed_url, err := url.Parse(img_url)
		must(err)
		urlPath := parsed_url.Path
		parts := strings.Split(urlPath, "/")
		file_name := parts[len(parts)-1]
		if file_name == "" {
			continue
		}
		img_url = fixupURL(img_url)
		imgPath := filepath.Join(imagesDir, file_name)
		writeURLToFile(img_url, imgPath)
		setAttr(img_tag, "src", "../../../images/"+file_name)
	}
	html, err := soup.Html()
	must(err)
	html = strings.Replace(html, "<html><head></head><body>", "", -1)
	html = strings.Replace(html, "</body></html>", "", -1)
	return html
}

// Function that creates the text describing the individual posts in a topic
func post_row(post_json *Post) string {
	avatar_url := post_json.AvatarTemplate
	parsed_url, err := url.Parse(avatar_url)
	must(err)

	parts := strings.Split(parsed_url.Path, "/")
	avatar_file_name := parts[len(parts)-1]
	avatar_url = fixupURL(avatar_url)
	avatar_url = strings.ReplaceAll(avatar_url, "{size}", "45")
	avatar_path := filepath.Join(imagesDir, avatar_file_name)
	writeURLToFileMust(avatar_url, avatar_path)

	user_name := post_json.Username
	html := postBodyTransform(post_json.Cooked)

	post_string := `      <div class="post_container">` + "\n"
	post_string += `        <div class="avatar_container">` + "\n"
	post_string += `          <img src="../../../images/` +
		avatar_file_name + `" class="avatar" />` + "\n"
	post_string += `        </div>` + "\n"
	post_string += `        <div class="post">` + "\n"
	post_string += `          <div class="user_name">` + user_name + `</div>` + "\n"
	post_string += `          <div class="post_content"` + ">\n"
	post_string += html + "\n"
	post_string += `          </div>` + "\n"
	post_string += `        </div>` + "\n"
	post_string += `      </div>` + "\n\n"
	return post_string
}

// The topic_row function generates the HTML for each topic on the main page
var category_id_to_name = map[int]string{}

func build_categories() {
	logf("build_categories\n")
	panicIf(len(category_id_to_name) > 0)
	var category_json CategoriesResponse
	var category_url = base_url + "/categories.json"
	// logf("category_url: %s\n", category_url)
	requests_get_json_must(category_url, &category_json)
	for _, cat := range category_json.CategoryList.Categories {
		id := cat.ID
		name := cat.Name
		category_id_to_name[id] = name
		logf("cat: %d => %s\n", id, name)
	}
}

func topic_row(topic_json *Topic) string {
	topic_html := `      <div class="topic-row">` + "\n"
	topic_url := "t/" + topic_json.Slug + "/" + strconv.Itoa(topic_json.ID)
	topic_title_text := topic_json.FancyTitle
	topic_post_count := topic_json.PostsCount
	topic_pinned := topic_json.PinnedGlobally
	topic_category := category_id_to_name[topic_json.CategoryID]

	topic_html += `        <span class="topic">`
	if topic_pinned {
		topic_html += `<i class="fa fa-thumb-tack"`
		topic_html += ` title="This was a pinned topic so it `
		topic_html += `appears near the top of the page."></i>`
	}
	topic_html += `<a href="` + topic_url + `">`
	topic_html += topic_title_text + `</a></span>` + "\n"
	topic_html += `        <span class="category">`
	topic_html += topic_category + `</span>` + "\n"
	topic_html += `        <span class="post-count">`
	topic_html += strconv.Itoa(topic_post_count) + `</span>` + "\n"
	topic_html += `      </div>` + "\n\n"
	return topic_html
}

// extract some information about site from HTML
func extract_site_info() {
	content := requests_get_must(base_url)
	r := bytes.NewBuffer(content)
	soup, err := goquery.NewDocumentFromReader(r)
	must(err)
	site_title = soup.Find("title").First().Text()
	logf("site_title: '%s'\n", site_title)

	siteLogoNode := soup.Find("img#site-logo")
	siteLogoURL, ok := siteLogoNode.First().Attr("src")
	logf("siteLogoURL: '%s'\n", siteLogoURL)
	if ok {
		// TODO:should use the right extension and update main_template
		//dst := filepath.Join(imagesDir, "site-logo"+filepath.Ext(siteLogoURL))
		dst := filepath.Join(imagesDir, "site-logo.png")
		writeURLToFileMust(siteLogoURL, dst)
	} else {
		dst := filepath.Join(imagesDir, "site-logo.png")
		writeFileMust(dst, site_logo_png)
	}
	{
		dst := filepath.Join(imagesDir, "missing_image.png")
		writeFileMust(dst, missing_image_png)
	}
}

func main() {
	if false {
		testPostBodyTransform()
		return
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: mirror-discourse <discourse-forum-url>\n")
		os.Exit(0)
	}
	base_url = strings.TrimSuffix(os.Args[1], "/")
	logf("base_url: '%s'\n", base_url)
	archive_blurb = "A partial archive of meta.discourse.org as of " + time.Now().String() + "."
	// TODO: format current date
	//+ date.today().strftime("%A %B %d, %Y") + '.'
	archive_blurb = ""

	// Templates for the webpages
	parsedURL, err := url.Parse(base_url)
	must(err)
	base_scheme = parsedURL.Scheme

	// The action is just starting here.

	// Check for the directory where plan to store things.
	// Note that this will be overwritten!
	must(os.RemoveAll(dstDir))
	must(os.MkdirAll(imagesDir, 0755))

	extract_site_info()

	// This is where *most* of the action happens.
	// The following bit of code grabs discourse_url/latest.json to generate a list of topics.
	// For each of these topics, we apply topic_row to generate a line on the main page.
	// If 'more_topics_url' appears in the response, we get more.

	// Note that there might be errors but the code does attempt to deal with them gracefully by
	// passing over them and continuing.
	//
	// My archive of DiscoureMeta generated 19 errors - all image downloads that replaced with a missing image PNG.

	build_categories()

	maxPages := 999
	if true {
		maxPages = 1
	}
	pageNo := 0
	topic_list_string := ""
	for pageNo < maxPages {
		uri := base_url + "/latest.json?no_definitions=true&page=" + fmt.Sprintf("%d", pageNo)
		pageNo++
		var topics TopicsResponse
		requests_get_json_must(uri, &topics)
		topic_list := topics.TopicList.Topics
		for _, topic := range topic_list {
			// logf("Topic: %#v\n", topic)
			write_topic(topic)
			topic_list_string += topic_row(topic)
		}
		uri_part := topics.TopicList.MoreTopicsURL
		if uri_part == "" {
			break
		}
	}

	file_string := main_template
	file_string = strings.ReplaceAll(file_string, "<!-- BANNER_HTML -->", banner_html)
	file_string = strings.Replace(file_string, "<!-- TITLE -->", site_title, -1)
	file_string = strings.Replace(file_string, "<!-- JUST_SITE_TITLE -->", site_title, -1)
	file_string = strings.Replace(file_string, "<!-- ARCHIVE_BLURB -->", archive_blurb, -1)
	file_string = strings.Replace(file_string, "<!-- TOPIC_LIST -->", topic_list_string, -1)

	{
		dst := filepath.Join(dstDir, "index.html")
		writeFileMust(dst, []byte(file_string))
	}

	{
		dst := filepath.Join(dstDir, "archived.css")
		writeFileMust(dst, []byte(css))
	}

	logf("Wrote website copy to %s\n", dstDir)
}