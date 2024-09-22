package main

import (
	"fmt"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// file, err := os.Open("/data/phishing_detection/phishy/6cd5cebf619dffacf08c46de33709e7c38beff2336e64394db36a22ba60eb26d/6cd5cebf619dffacf08c46de33709e7c38beff2336e64394db36a22ba60eb26d.html")
	file, err := os.Open("/data/phishing_detection/benign/aafc21845bccf1ad2181e9dd53b5cf5b1db029ed5a95bf106b4d7c9670b7e0f3/aafc21845bccf1ad2181e9dd53b5cf5b1db029ed5a95bf106b4d7c9670b7e0f3.html")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	baseURL := "https://learn.microsoft.com/en-us/sysinternals/#123"
	// TODO: Process the origin # URL

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL")
	}

	baseURLScheme := parsedBaseURL.Scheme
	baseURLHost := parsedBaseURL.Host
	basedURLPath := parsedBaseURL.Path
	fmt.Println("basedURLPath: ", basedURLPath)
	processedBaseURL := baseURLScheme + "://" + baseURLHost + basedURLPath

	links := extractAndFormatLinks(doc, processedBaseURL, parsedBaseURL)

	// var waitForCrawlUrls []string
	// for _, link := range links {
	// 	if link != baseURL {
	// 		parsedURL, err := url.Parse(baseURL)
	// 		if err != nil {
	// 			fmt.Println("Error parsing URL")
	// 		}
	// 		scheme := parsedURL.Scheme
	// 		host := parsedURL.Host
	// 		if host != baseURLHost {
	// 			waitForCrawlUrls = append(waitForCrawlUrls, link)
	// 		} else {
	// 		}
	//
	// 	}
	// }

	fmt.Print(links)
}
