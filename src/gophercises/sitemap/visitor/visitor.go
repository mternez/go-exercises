package visitor

import (
	"net/http"
	"sitemap/parser"

	"golang.org/x/net/html"
)

func Visit(link *parser.Link, host *string) {

	alreadyVisitedInternalLinks := make(map[string]*parser.Link)
	alreadyVisitedInternalLinks[link.Href] = link
	VisitInternalLinks(link, host, alreadyVisitedInternalLinks)
}

func VisitInternalLinks(link *parser.Link, host *string, alreadyVisitedInternalLinks map[string]*parser.Link) {

	if !link.Internal {
		return
	}

	// TODO : handle case where host doesn't end with "/"
	resp, err := http.Get(*host + link.Href)

	if err != nil {
		panic("Couldn't open link")
	}

	alreadyVisitedInternalLinks[link.Href] = link

	page, _ := html.Parse(resp.Body)

	linksMap := parser.GetLinks(page, *host)
	link.Children = linksMap

	for key, value := range linksMap {
		_, ok := alreadyVisitedInternalLinks[key]
		if !ok && value.Internal {
			alreadyVisitedInternalLinks[key] = value
			VisitInternalLinks(value, host, alreadyVisitedInternalLinks)
		}
	}
}
