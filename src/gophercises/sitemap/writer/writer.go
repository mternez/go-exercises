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

func (writer *XMLLinkWriter) Write(link *parser.Link) {

	marshalledBytes, err := xml.MarshalIndent(link, " ", " ")

	if err != nil {
		panic("Failed to marshal link : " + link.Href + " - " + err.Error())
	}

	numberBytesWritten, err := writer.file.Write(marshalledBytes)

	if numberBytesWritten <= 0 || err != nil {
		panic("Failed to write marshalled link to file : " + link.Href + " - " + err.Error())
	}
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
