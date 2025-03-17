package main

import (
	"net/http"
	"sitemap/parser"
	"sitemap/visitor"

	"golang.org/x/net/html"
)

func main() {

	// /!\ Host should end with "/"
	// Or handle the case described in TODO in visitor.go
	host := "https://www.calhoun.io/"

	resp, err := http.Get(host)

	if err != nil {
		panic("Couldn't open link")
	}

	mainPage, err := html.Parse(resp.Body)

	if err != nil {
		panic("Couldn't read response body")
	}

	linksMap := parser.GetLinks(mainPage, host)

	for _, value := range linksMap {
		visitor.Visit(value, &host)
	}
}
