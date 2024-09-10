package main

import (
	"chromedp-crawler/crawler"
	"log"
)

func main() {
	url := "https://www.google.com"

	links, err := crawler.CollectLinks(url)
	if err != nil {
		log.Fatalf("Error collecting links: %v", err)
	}
	log.Printf("Collected links: ", links)
}
