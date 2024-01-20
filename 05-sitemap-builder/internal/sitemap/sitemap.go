package sitemap

import (
	"05-sitemap-builder/internal/pkg/links"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

type loc struct {
	Value string `xml:"loc"`
}

func Build(urlString string, maxDepth int) ([]byte, error) {
	normalizedUrl := normalized(urlString)
	urls := bfs(normalizedUrl, maxDepth)
	return encodeXml(urls)
}

func encodeXml(urls []string) ([]byte, error) {
	toEncode := urlset{
		Xmlns: xmlns,
	}

	for _, page := range urls {
		toEncode.Urls = append(toEncode.Urls, loc{page})
	}

	var buf bytes.Buffer

	buf.WriteString(xml.Header)

	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	err := enc.Encode(toEncode)
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(&buf)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func bfs(urlString string, maxDepth int) []string {
	visited := make(map[string]struct{})

	var queue map[string]struct{}
	nextQueue := map[string]struct{}{
		urlString: {},
	}

	var wg sync.WaitGroup

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})
		if len(queue) == 0 {
			break
		}

		hrefChan := make(chan []string)

		for queuedUrl := range queue {
			if _, ok := visited[queuedUrl]; ok {
				continue
			}
			visited[queuedUrl] = struct{}{}

			wg.Add(1)
			go worker(queuedUrl, hrefChan, &wg)

		}

		go func() {
			wg.Wait()
			close(hrefChan)
		}()

		for hrefs := range hrefChan {
			for _, href := range hrefs {
				if _, ok := visited[href]; !ok {
					nextQueue[href] = struct{}{}
				}
			}
		}
	}

	result := make([]string, 0, len(visited))
	for visitedUrl := range visited {
		result = append(result, visitedUrl)
	}

	return result
}

func worker(
	urlString string,
	hrefChan chan []string,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	hrefs := getValidHrefs(urlString)

	fmt.Println("done:", urlString)

	hrefChan <- hrefs
}

func getValidHrefs(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	r := bytes.NewReader(body)
	siteLinks, err := links.Parse(r)
	if err != nil {
		return nil
	}

	domainUrl := getDomainUrl(resp.Request.URL)

	var result []string

	for _, link := range siteLinks {
		if strings.HasPrefix(link.Href, "/") {
			link.Href = domainUrl + link.Href
		}
		if strings.HasPrefix(link.Href, domainUrl) {
			normalizedLink := normalized(link.Href)
			if !slices.Contains(result, normalizedLink) {
				result = append(result, normalizedLink)
			}
		}
	}

	return result
}

func normalized(urlString string) string {
	if strings.Contains(urlString, "?") {
		urlString = strings.Split(urlString, "?")[0]
	}
	if strings.Contains(urlString, "#") {
		urlString = strings.Split(urlString, "#")[0]
	}
	return strings.TrimSuffix(urlString, "/")
}

func getDomainUrl(urlStruct *url.URL) string {
	protocol := urlStruct.Scheme
	host := urlStruct.Host

	return fmt.Sprintf("%v://%v", protocol, host)
}
