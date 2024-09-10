package crawler

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func CollectLinks(url string) (map[string]struct{}, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var contentType string
	var links []string

	err1 := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// chromedp.WaitVisible(`body > footer`),
		chromedp.Sleep(10*time.Second),
		chromedp.Evaluate(`document.contentType`, &contentType),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),
	)

	if err1 != nil {
		return nil, err1
	}

	if contentType != "text/html" {
		log.Printf("URL is not an HTML page: %s (content-type: %s)", url, contentType)
		return nil, nil
	}

	linkSet := make(map[string]struct{})
	for _, link := range links {
		linkSet[link] = struct{}{}
	}

	return linkSet, nil
}
