package main

import (
	"chromedp-crawler/crawler"
	"chromedp-crawler/storage"
	"log"
)

func main() {
	test_url := "https://www.bilibili.com"

	log.Printf(`Crawling target url: [%s]`, test_url)
	contentType, links, error := crawler.CollectLinks(test_url)
	if error != nil {
		log.Printf("Error collecting links: %v", error)
	}
	log.Printf(`Collected links: %s`, links)
	if contentType == "text/html" {
		storage.SaveLinks(test_url, links)
	}

	// for _, link := range links {
	//
	// 	log.Printf("Crawling link: [%s]", link)
	// 	time.Sleep(1 * time.Second)
	// 	contentType, links, error := crawler.CollectLinks(link)
	// 	if error != nil {
	// 		log.Print(error)
	// 		continue
	// 	}
	// 	log.Printf(`Links for [%s]: %s`, link, links)
	// 	if contentType == "text/html" {
	// 		storage.SaveLinks(link, links)
	// 	}
	// }
}
