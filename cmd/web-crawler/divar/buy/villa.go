package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/chromedp/chromedp"
)

// Site represents a site to scrape with a starting URL and a selector for links
type Site struct {
	BaseURL       string
	LinkSelector  string
	TitleSelector string
}

func main() {
	// Define the sites to scrape
	sites := []Site{
		{
			BaseURL:       "https://divar.ir/s/iran/buy-villa",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "h1.kt-page-title__title.kt-page-title__title--responsive-sized",
		},
		// Add more sites as needed
	}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)
		go func(site Site) {
			defer wg.Done()
			scrapeSite(site)
		}(site)
	}

	// Wait for all sites to finish scraping
	wg.Wait()
	fmt.Println("Scraping completed for all sites.")
}

func scrapeSite(site Site) {
	// Create a new ChromeDP context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Store the page titles for this site
	var pageTitles []string

	// Load the base URL
	err := chromedp.Run(ctx,
		chromedp.Navigate(site.BaseURL),
		chromedp.WaitVisible(site.LinkSelector, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Printf("Failed to load base URL %s: %v", site.BaseURL, err)
		return
	}

	// Extract all links with the specified CSS class
	var links []string
	err = chromedp.Run(ctx,
		chromedp.Evaluate(fmt.Sprintf(`Array.from(document.querySelectorAll("%s")).map(a => a.href)`, site.LinkSelector), &links),
	)
	if err != nil {
		log.Printf("Failed to retrieve links from %s: %v", site.BaseURL, err)
		return
	}

	// Visit each link and extract the specified title
	for _, link := range links {
		var pageTitle string
		err = chromedp.Run(ctx,
			chromedp.Navigate(link),
			chromedp.WaitVisible(site.TitleSelector, chromedp.ByQuery),
			chromedp.Text(site.TitleSelector, &pageTitle, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("Failed to scrape page %s: %v", link, err)
			continue
		}
		pageTitles = append(pageTitles, pageTitle)

		// Optional: Return to the base URL after each link (if necessary)
		err = chromedp.Run(ctx, chromedp.Navigate(site.BaseURL))
		if err != nil {
			log.Printf("Failed to return to base URL %s: %v", site.BaseURL, err)
		}
	}

	// Output the collected page titles for this site
	fmt.Printf("Titles for %s:\n", site.BaseURL)
	for i, title := range pageTitles {
		fmt.Printf("Page %d title: %s\n", i+1, title)
	}
}
