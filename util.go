package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/kjk/common/u"
	"golang.org/x/net/html"
)

var (
	panicIf = u.PanicIf
	must    = u.Must
)

func nodeGetAttr(attrs []html.Attribute, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}

func nodeGetAttrMust(attrs []html.Attribute, name string) string {
	s, ok := nodeGetAttr(attrs, name)
	if !ok {
		panic(fmt.Sprintf("didn't find attribute '%s'", name))
	}
	return s
}

func nodeSetAttr(node *html.Node, name string, val string) {
	for i, attr := range node.Attr {
		if attr.Key == name {
			node.Attr[i].Val = val
			return
		}
	}
	panic(fmt.Sprintf("didn't find attribute '%s'", name))
}

func logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func httpGetRetry(uri string, delay time.Duration) ([]byte, error) {
	if delay != 0 {
		time.Sleep(delay)
	}
	rsp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == 429 {
		switch delay {
		case 0:
			delay = time.Second
		case time.Second:
			delay = time.Second * 3
		default:
			return nil, fmt.Errorf("failed with status code %d", rsp.StatusCode)
		}
		return httpGetRetry(uri, delay)
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("failed with status code %d", rsp.StatusCode)
	}
	return ioutil.ReadAll(rsp.Body)
}

func httpGet(uri string) ([]byte, error) {
	logf("httpGet: '%s'\n", uri)
	return httpGetRetry(uri, 0)
}

func httpGetCached(uri string, cacheDir string) ([]byte, error) {
	if flgNoCache {
		return httpGet(uri)
	}
	var ext string
	pu, err := url.Parse(uri)
	if err == nil {
		ext = filepath.Ext(pu.Path)
	}

	fileName := u.DataSha1Hex([]byte(uri)) + ext
	path := filepath.Join(cacheDir, fileName)
	d, err := os.ReadFile(path)
	if err == nil {
		logf("Read '%s' from '%s'\n", uri, path)
		return d, nil
	}
	d, err = httpGet(uri)
	if err != nil {
		return nil, err
	}
	// it's ok if we fail to write to cache
	err = os.MkdirAll(cacheDir, 0755)
	must(err)
	err = os.WriteFile(path, d, 0644)
	must(err)
	logf("Wrote '%s' to '%s'\n", uri, path)
	return d, nil
}

// func httpGetCachedMust(uri string, cacheDir string) []byte {
// 	d, err := httpGetCached(uri, cacheDir)
// 	must(err)
// 	return d
// }

func httpGetMust(uri string) []byte {
	d, err := httpGet(uri)
	must(err)
	return d
}

// func httpGetJSONMust(uri string, v interface{}) {
// 	d, err := httpGet(uri)
// 	must(err)
// 	err = json.Unmarshal(d, v)
// 	must(err)
// }

func httpGetJSONCachedMust(uri string, v interface{}, cacheDir string) {
	d, err := httpGetCached(uri, cacheDir)
	must(err)
	err = json.Unmarshal(d, v)
	must(err)
}

func fixupURL(uri string) string {
	parsed_url, err := url.Parse(uri)
	must(err)
	if base_scheme == "" {
		base_scheme = "https"
	}
	if parsed_url.Host == "" {
		res := base_url + uri
		// logf("fixupURL: %s => %s\n", uri, res)
		return res
	}
	if parsed_url.Scheme == "" {
		res := base_scheme + ":" + uri
		// logf("fixupURL: %s => %s\n", uri, res)
		return res
	}
	return uri
}

func writeFileMust(path string, d []byte) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	must(err)
	err = os.WriteFile(path, d, 0644)
	must(err)
	logf("Wrote %s\n", path)
}

func writeFile(path string, d []byte) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, d, 0644)
	if err != nil {
		return err
	}
	logf("Wrote %s\n", path)
	return nil
}
