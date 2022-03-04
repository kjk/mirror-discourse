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

func getAttr(attrs []html.Attribute, name string) (string, bool) {
	for _, attr := range attrs {
		if attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}

func getAttrMust(attrs []html.Attribute, name string) string {
	s, ok := getAttr(attrs, name)
	if !ok {
		panic(fmt.Sprintf("didn't find attribute '%s'", name))
	}
	return s
}

func setAttr(node *html.Node, name string, val string) {
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

func requests_get_with_delay(uri string, delay time.Duration) ([]byte, error) {
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
		return requests_get_with_delay(uri, delay)
	}
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("failed with status code %d", rsp.StatusCode)
	}
	return ioutil.ReadAll(rsp.Body)
}

func requests_get(uri string) ([]byte, error) {
	logf("request_get: '%s'\n", uri)
	return requests_get_with_delay(uri, 0)
}

func requests_get_must(uri string) []byte {
	d, err := requests_get(uri)
	must(err)
	return d
}

func requests_get_json_must(uri string, v interface{}) {
	logf("requests_get_json_must: %s\n", uri)
	d, err := requests_get(uri)
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
		logf("fixupURL: %s => %s\n", uri, res)
		return res
	}
	if parsed_url.Scheme == "" {
		res := base_scheme + ":" + uri
		logf("fixupURL: %s => %s\n", uri, res)
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
