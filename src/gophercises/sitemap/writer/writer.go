package writer

import (
	"encoding/xml"
	"os"
	"sitemap/parser"
)

type XMLLinkWriter struct {
	file          *os.File
	headerWritten bool
}

type UrlSet struct {
	XMLName xml.Name       `xml:"urlset"`
	Xmlns   string         `xml:"xmlns,attr"`
	Urls    []*parser.Link `xml:"url"`
}

func (writer *XMLLinkWriter) Write(links []*parser.Link) {

	urlSet := &UrlSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9", Urls: links}

	marshalledBytes, err := xml.MarshalIndent(urlSet, " ", " ")

	if err != nil {
		panic("Failed to marshal link : " + err.Error())
	}

	numberBytesWritten, err := writer.file.Write(marshalledBytes)

	if numberBytesWritten <= 0 || err != nil {
		panic("Failed to write marshalled link to file : " + err.Error())
	}
}

func (writer *XMLLinkWriter) Close() error {
	return writer.file.Close()
}

func NewXMLLinkWriter(filePath string) *XMLLinkWriter {

	file, err := os.Open(filePath)

	if err != nil {
		file, err = os.Create(filePath)
	}

	if err != nil {
		panic("Couldn't open or create file at " + filePath + " - " + err.Error())
	}

	file.Write([]byte(xml.Header))

	return &XMLLinkWriter{file: file, headerWritten: true}
}
