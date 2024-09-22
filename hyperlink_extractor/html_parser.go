package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func extractAndFormatLinks(n *html.Node, baseURL string, parsedURL *url.URL) []string {
	var links []string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {

				link := attr.Val

				// Absolute paths and relative paths are handled in two cases
				if strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "http://") {
					links = append(links, link)
				} else {
					scheme := parsedURL.Scheme
					host := parsedURL.Host

					if strings.HasPrefix(link, "#") {
						link = baseURL + link
						links = append(links, link)
					} else if strings.HasPrefix(link, "//") {
						link = scheme + ":" + link
						links = append(links, link)
					} else if strings.HasPrefix(link, "/") {
						link = scheme + "://" + host + link
						links = append(links, link)
					} else if strings.HasPrefix(link, "..") {
						fmt.Printf("Relative path not supported: %s", link)
					} else if !strings.Contains(link, ":") {
						link = baseURL + link
						links = append(links, link)
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractAndFormatLinks(c, baseURL, parsedURL)...)
	}

	removedDuplicatesLinks := RemoveDuplicates(links)

	return removedDuplicatesLinks
}
