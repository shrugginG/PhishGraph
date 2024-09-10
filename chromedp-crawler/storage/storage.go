package storage

import (
	"chromedp-crawler/utils"
	"context"
	"encoding/json"
	"log"
)

func SaveLinks(url string, links map[string]struct{}) {
	log.Printf("Saving %d unique links", len(links))
	for link := range links {
		log.Print(link)
	}

	conn, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect datasbase", err)
	}

	linksList := make([]string, 0, len(links))
	for link := range links {
		linksList = append(linksList, link)
	}

	linksJSON, err := json.Marshal(linksList)
	if err != nil {
		log.Fatalf("Error marshalling links: %v", err)
	}

	query := `INSERT INTO webpage_links (url, links) VALUES ($1, $2)`
	_, err = conn.Exec(context.Background(), query, url, linksJSON)
	if err != nil {
		log.Fatalf("Failed to insert links: %v", err)
	}

	log.Println("Successfully saved links to database")
}
