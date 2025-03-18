package parser

import (
	"encoding/xml"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	XMLName     xml.Name         `xml:"url"`
	Href        string           `xml:"-"`
	Children    map[string]*Link `xml:"-"`
	TrimmedHref string           `xml:"loc"`
	Internal    bool             `xml:"-"`
}

func GetLinks(node *html.Node, host string) map[string]*Link {

	links := make(map[string]*Link)

	for child := range node.Descendants() {
		if child.Type == html.ElementNode && child.Data == "a" {
			var href string = getAttribute(child, "href")
			if href != "" {
				links[href] = newLink(host, href)
			}
		}
	}

	return links
}

func newLink(host string, href string) *Link {

	var isLinkInternalRegExp = regexp.MustCompile("^(/)|" + "^(" + host + ")")

	if isLinkInternalRegExp.Match([]byte(href)) {
		trimmedHref := trimHref(href, host)
		return &Link{Href: href, TrimmedHref: trimmedHref, Internal: true}
	} else {
		return &Link{Href: href, Internal: false}
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
