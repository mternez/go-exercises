package parser

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href        string
	Children    map[string]*Link
	TrimmedHref string
	Internal    bool
	host        string
}

func GetLinks(node *html.Node, host string) map[string]*Link {

	links := make(map[string]*Link)

	for child := range node.Descendants() {
		if child.Type == html.ElementNode && child.Data == "a" {
			var href string = getAttribute(child, "href")
			if href != "" {
				links[href] = NewLink(host, href)
			}
		}
	}

	return links
}

func NewLink(host string, href string) *Link {

	var isLinkInternalRegExp = regexp.MustCompile("^(/)|" + "^(" + host + ")")

	if isLinkInternalRegExp.Match([]byte(href)) {
		trimmedHref := trimHref(href, host)
		return &Link{Href: href, TrimmedHref: trimmedHref, Internal: true, host: host}
	} else {
		return &Link{Href: href, Internal: false, host: host}
	}
}

func getAttribute(node *html.Node, key string) string {

	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

func trimHref(href string, host string) string {
	trimmedHref := strings.Replace(href, host, "", 1)
	if trimmedHref != "" {
		if trimmedHref[0] != '/' {
			trimmedHref = "/" + trimmedHref
		}
	}
	return trimmedHref
}
