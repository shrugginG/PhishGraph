package storage

import (
	"chromedp-crawler/utils"
	"context"
	"encoding/json"
	"log"
)

func SaveLinks(url string, links []string) {
	// log.Printf("Saving %d unique links", len(links))

	conn, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf(`Unable to connect datasbase: %v`, err)
	}

	linksJSON, err := json.Marshal(links)
	if err != nil {
		log.Fatalf("Error marshalling links: %v", err)
	}

	query := `INSERT INTO webpage_links (url, links) VALUES ($1, $2) ON CONFLICT (url) DO NOTHING`
	_, err = conn.Exec(context.Background(), query, url, linksJSON)
	if err != nil {
		log.Fatalf("Failed to insert links: %v", err)
	}

	log.Println("Successfully saved links to database")
}
