package main

import (
	"sitemap/visitor"
	"sitemap/writer"
)

func main() {

	// /!\ Host should end with "/"
	// Or handle the case described in TODO in visitor.go
	host := "https://www.calhoun.io/"

	xmlWriter := writer.NewXMLLinkWriter("./sitemap.xml")
	linksMap := visitor.Visit(host)
	xmlWriter.Write(linksMap.GetValues())
	xmlWriter.Close()
}
