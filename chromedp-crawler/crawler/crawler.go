package crawler

import (
	"chromedp-crawler/utils"
	"context"

	// "fmt"
	"log"
	"time"

	// "github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func CollectLinks(url string) (string, []string, error) {
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	//
	// ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	// defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36')"),
		chromedp.IgnoreCertErrors,

		chromedp.Flag("headless", false),
		chromedp.Flag("diable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.ProxyServer("socks5://127.0.0.1:7898"),
		chromedp.Flag("auto-open-devtools-for-tabs", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// var statusCode int64
	// var firstRequestID network.RequestID

	// chromedp.ListenTarget(ctx, func(ev interface{}) {
	// 	switch ev := ev.(type) {
	// 	case *network.EventRequestWillBeSent:
	// 		log.Printf(`Request URL: %s`, ev.Request.URL)
	// 		if firstRequestID == "" {
	// 			firstRequestID = ev.RequestID
	// 		}
	//
	// 	case *network.EventResponseReceived:
	// 		fmt.Print(ev.Response.Status)
	// 		if ev.RequestID == firstRequestID {
	// 			statusCode = ev.Response.Status
	// 		}
	// 	}
	// })

	var contentType string
	var links []string

	err1 := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// chromedp.WaitVisible(`body > footer`),
		// chromedp.WaitReady(`document.readyState === 'complete'`),
		chromedp.Sleep(15*time.Second),
		chromedp.Evaluate(`document.contentType`, &contentType),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),
	)

	log.Printf(`[%s] content-type: %s`, url, contentType)

	if err1 != nil {
		return "", nil, err1
	}

	if contentType != "text/html" {
		// log.Printf(`URL status code: %d`, statusCode)
		log.Printf(`URL is not an HTML page: %s (content-type: %s)`, url, contentType)
	}

	removedDuplicateLinks := utils.RemoveDuplicates(links)

	// linkSet := make(map[string]struct{})
	// for _, link := range links {
	// 	linkSet[link] = struct{}{}
	// }
	// log.Printf(`URL status code: %d`, statusCode)

	return contentType, removedDuplicateLinks, nil
}
