package main

import (
	"chromedp-crawler/crawler"
	"chromedp-crawler/storage"
	"fmt"
	"log"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// file, err := os.Open("/data/phishing_detection/phishy/6cd5cebf619dffacf08c46de33709e7c38beff2336e64394db36a22ba60eb26d/6cd5cebf619dffacf08c46de33709e7c38beff2336e64394db36a22ba60eb26d.html")
	// file, err := os.Open("/home/jxlu/project/PhishDetect/PhishGraph/data/aafc21845bccf1ad2181e9dd53b5cf5b1db029ed5a95bf106b4d7c9670b7e0f3/aafc21845bccf1ad2181e9dd53b5cf5b1db029ed5a95bf106b4d7c9670b7e0f3.html")
	basePath := "/home/jxlu/project/PhishDetect/PhishGraph/data"
	urlsha256 := "448ad8d45e8b8b9559610161b9f08484390b04cf1540c73aff1cb01484c70d8f"
	htmlFilePath := fmt.Sprintf("%s/%s/%s.html", basePath, urlsha256, urlsha256)

	file, err := os.Open(htmlFilePath)
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

	// baseURL := "https://learn.microsoft.com/en-us/sysinternals/"
	baseURL := "https://www.botanical-journeys-plant-guides.com/"
	// TODO: Process the origin # URL

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL")
	}

	baseURLScheme := parsedBaseURL.Scheme
	baseURLHost := parsedBaseURL.Host
	basedURLPath := parsedBaseURL.Path
	fmt.Println("basedURLPath: ", basedURLPath)
	fmt.Println("basedURLFrag: ", parsedBaseURL.Fragment)

	processedBaseURL := baseURLScheme + "://" + baseURLHost + basedURLPath

	links := crawler.ExtractAndFormatLinks(doc, processedBaseURL, parsedBaseURL)

	// linksJsonFile := "/home/jxlu/project/PhishDetect/PhishGraph/data/aafc21845bccf1ad2181e9dd53b5cf5b1db029ed5a95bf106b4d7c9670b7e0f3/links.json"
	linksJsonFile := fmt.Sprintf("%s/%s/links.json", basePath, urlsha256)
	storage.SaveLinksAsJson(linksJsonFile, links)

	var waitForCrawlUrls []string
	for _, link := range links {
		parsedURL, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing URL")
		}

		if parsedURL.Host != baseURLHost {
			waitForCrawlUrls = append(waitForCrawlUrls, link)
		} else {
			if parsedURL.Path != basedURLPath {
				waitForCrawlUrls = append(waitForCrawlUrls, link)
			}
		}
	}

	fmt.Println(links)
	fmt.Println(len(links))
	fmt.Println(waitForCrawlUrls)
	fmt.Println(len(waitForCrawlUrls))

	// test_url := "https://learn.microsoft.com
	//
	// log.Printf(`Crawling target url: [%s]`, test_url)
	// contentType, links, error := crawler.CollectLinks(test_url)
	// if error != nil {
	// 	log.Printf("Error collecting links: %v", error)
	// }
	// log.Printf(`Collected links: %s`, links)
	// if contentType == "text/html" {
	// 	fmt.Print("Hello")
	// 	storage.SaveLinks(test_url, links)
	// }
	//
	for _, link := range waitForCrawlUrls {

		log.Printf("Crawling link: [%s]", link)
		// time.Sleep(1 * time.Second)
		contentType, links, error := crawler.CollectLinks(link)
		if error != nil {
			log.Print(error)
			continue
		}
		log.Printf(`Links for [%s]: %s`, link, links)
		if contentType == "text/html" {
			storage.SaveLinks(link, links)
		}
	}
}
