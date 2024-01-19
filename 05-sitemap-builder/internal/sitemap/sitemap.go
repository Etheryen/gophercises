package sitemap

import (
	"05-sitemap-builder/internal/pkg/links"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

type SiteNode struct {
	Location string
	Children []SiteNode
}

func Build(urlString string) (SiteNode, error) {
	normalizedUrl := normalized(urlString)
	visited := []string{normalizedUrl}
	return getSiteNode(normalizedUrl, visited)
}

func normalized(urlString string) string {
	if !strings.HasSuffix(urlString, "/") {
		return urlString + "/"
	}
	return urlString
}

func getSiteNode(urlString string, visited []string) (SiteNode, error) {
	validHrefs, err := getValidHrefs(urlString)
	if err != nil {
		return SiteNode{}, err
	}

	fmt.Println("\nLocation:", urlString)
	fmt.Printf("Children: %v\n\n", validHrefs)

	var childrenNodes []SiteNode

	for _, href := range validHrefs {
		if !slices.Contains(visited, href) {
			visited = append(visited, href)
			childNode, err := getSiteNode(href, visited)
			if err != nil {
				return SiteNode{}, err
			}
			childrenNodes = append(childrenNodes, childNode)
		} else {
			childrenNodes = append(childrenNodes, SiteNode{href, nil})
		}
	}

	node := SiteNode{
		Location: urlString,
		Children: childrenNodes,
	}

	return node, nil
}

// TODO: remove foreign domain links, add domain to relative ones,
// 		 maybe normalize the rest (add protocol or sth)

func getValidHrefs(urlString string) ([]string, error) {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	siteLinks, err := links.Parse(r)
	if err != nil {
		return nil, err
	}

	domainUrl := getDomainUrl(resp.Request.URL)

	var result []string

	for _, link := range siteLinks {
		// FIXME: make it so that links to HOME
		//		  (e.g. / or domainUrl) still get added
		if link.Href == urlString ||
			(link.Href == "/" && urlString == domainUrl) {
			continue
		}
		if strings.HasPrefix(link.Href, "/") {
			link.Href = domainUrl + link.Href[1:]
		}
		if strings.HasPrefix(link.Href, domainUrl) {
			linkWithoutQuery := cleaned(link.Href)
			if !slices.Contains(result, linkWithoutQuery) {
				result = append(result, linkWithoutQuery)
			}
		}
	}

	return result, nil
}

func cleaned(urlString string) string {
	if strings.Contains(urlString, "?") {
		return strings.Split(urlString, "?")[0]
	}
	if strings.Contains(urlString, "#") {
		return strings.Split(urlString, "#")[0]
	}
	return urlString
}

func getDomainUrl(urlStruct *url.URL) string {
	protocol := urlStruct.Scheme
	host := urlStruct.Host

	return fmt.Sprintf("%v://%v/", protocol, host)
}
