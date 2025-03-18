package main

import (
	"sitemap/parser"
	"sitemap/visitor"
	"sitemap/writer"
)

func main() {

	// /!\ Host should end with "/"
	// Or handle the case described in TODO in visitor.go
	host := "https://www.calhoun.io/"

	xmlWriter := writer.NewXMLLinkWriter("./sitemap.xml")
	visitor.Visit(host).Range(func(key string, value *parser.Link) {
		xmlWriter.Write(value)
	})
}
