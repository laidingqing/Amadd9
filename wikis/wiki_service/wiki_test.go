package wiki_service_test

import (
	"net/url"
	"regexp"
	"testing"
)

func TestUsers(t *testing.T) {
	str := "window._d7666e17bbd38568_ = page_count:\"sdfwdf\", {\"scoreasdfasdfasdf\"source:\"https:\u002F\u002Fd12drcwhcokzqv.cloudfront.net\u002F3612676.gp4\","

	re, _ := regexp.Compile(`source:"[\s\S]*"`)
	source := re.FindString(str)
	urlStr := source[8 : len(source)-1]
	t.Log(urlStr)
	url, _ := url.PathUnescape(urlStr)

	t.Log(url)
}
