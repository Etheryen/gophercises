package parsing

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := searchLinks(doc)

	return links, nil
}

func searchLinks(n *html.Node) []Link {
	var links []Link

	if n.Type == html.ElementNode && n.Data == "a" {
		href := getHrefFromAttr(n.Attr)
		text := getTextFromNode(n.FirstChild)
		trimmed := strings.TrimSpace(text)

		links = append(links, Link{href, trimmed})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, searchLinks(c)...)
	}

	return links
}

func getTextFromNode(n *html.Node) string {
	if n == nil {
		return ""
	}

	result := ""

	if n.Type == html.TextNode {
		result += n.Data
	}

	result += getTextFromNode(n.FirstChild)
	result += getTextFromNode(n.NextSibling)

	return result
}

func getHrefFromAttr(attr []html.Attribute) string {
	for _, a := range attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}
