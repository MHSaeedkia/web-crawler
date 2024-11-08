package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// Define a struct to store advertisement data
type Advertisement struct {
	Title       string
	Price       string
	Description string
}

func main() {
	// Create a new context for chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Define slices to store scraped data
	var adsTitle []string
	var adsPrice []string
	var adsDescription []string

	// Slice to store advertisement details
	var ads []Advertisement

	// Run chromedp tasks
	err := chromedp.Run(ctx,
		// Navigate to the Divar site
		chromedp.Navigate("https://divar.ir/s/tehran"),

		// Wait for the page to load
		chromedp.Sleep(2*time.Second),

		// Extract advertisement details
		chromedp.Tasks{
			// Get all ad titles
			chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__title")).map(el => el.innerText)`, &adsTitle),
			// Get all ad prices
			chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__description")).map(el => el.innerText)`, &adsPrice),
			// Get all ad descriptions
			chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__bottom-description")).map(el => el.innerText)`, &adsDescription),
		},
	)

	if err != nil {
		log.Fatal("Error while scraping:", err)
	}

	// Append all ads data to ads slice
	for i := 0; i < len(adsTitle); i++ {
		ads = append(ads, Advertisement{
			Title:       adsTitle[i],
			Price:       adsPrice[i],
			Description: adsDescription[i],
		})
	}

	// Print out all scraped advertisements
	for _, ad := range ads {
		fmt.Printf("Title: %s\nPrice: %s\nDescription: %s\n\n", ad.Title, ad.Price, ad.Description)
	}
}
