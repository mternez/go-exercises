package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Garbage
func ParseLinks(filePath string) []Link {

	var err error
	file, err := os.Open(filePath)
	if err != nil {
		panic("Couldn't open file : " + filePath)
	}
	var tag string
	var content string
	links := make([]Link, 0)
	for {
		tag, err = findOpeningTag(file, "a")
		if err != nil {
			break
		}
		content, err = readUntilClosingTag(file, "a")
		if tag != "" && content != "" {
			links = append(links, Link{Href: tag, Text: strings.Trim(content, " ")})
		}
		if err != nil {
			break
		}
	}
	return links
}

// Garbage
func findOpeningTag(file *os.File, tag string) (string, error) {
	var err error
	var next string
	tagReadFully := false
	result := make([]string, 1)
	for {
		if tagReadFully {
			break
		}
		next, err = getNext(file)
		if err != nil {
			break
		}
		if next == "<" {
			next, err = getNext(file)
			if err != nil {
				break
			}
			if next == tag {
				result = append(result, "<")
				result = append(result, next)
				// Read rest of the tag
				for {
					next, err = getNext(file)
					if err != nil {
						break
					}
					result = append(result, next)
					if next == ">" {
						tagReadFully = true
						break
					}
				}
			}
		}
	}
	return strings.Join(result, ""), err
}

// Garbage
func readUntilClosingTag(file *os.File, tag string) (string, error) {
	var err error
	var next string
	tagReadFully := false
	result := make([]string, 1)
	for {
		if tagReadFully {
			break
		}
		next, err = getNext(file)
		if err != nil {
			break
		}
		if next == "<" {
			next, err = getNext(file)
			if err != nil {
				break
			}
			if next == "/" {
				next, err = getNext(file)
				if err != nil {
					break
				}
				if next == tag {
					tagReadFully = true
					break
				}
			}
		}
		result = append(result, next)
	}
	if tagReadFully {
		return strings.Join(result, ""), err
	}
	return "", errors.New("Closing tag not found")
}

// Garbage
func getNext(file *os.File) (string, error) {
	buffer := make([]byte, 1)
	_, err := io.ReadAtLeast(file, buffer, 1)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func ParseWithXHTML(filPath string) []Link {

	file, err := os.Open(filPath)

	node, err := html.Parse(file)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return nil
	}

	return getLinks(node)
}

func getLinks(node *html.Node) []Link {

	links := make([]Link, 1)

	for child := range node.Descendants() {
		if child.Type == html.ElementNode && child.Data == "a" {
			link := Link{}
			for _, attr := range child.Attr {
				if attr.Key == "href" {
					link.Href = attr.Val
				}
			}
			link.Text = getNodeContent(child)
			links = append(links, link)
		}
	}

	return links
}

func getNodeContent(node *html.Node) string {

	content := make([]string, 1)

	for descendant := range node.Descendants() {
		content = append(content, descendant.Data)
	}

	return strings.Join(content, "")
}
