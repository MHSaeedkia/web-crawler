package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// Define a struct to store advertisement data
type Advertisement struct {
	Title       string
	Price       string
	Description string
	Details     string   // Detailed description from the ad page
	Images      []string // URLs of the images from the ad page
}

func main() {
	// Define slices to store scraped data
	var adsTitle []string
	var adsPrice []string
	var adsDescription []string
	var adsLinks []string // Store links to each ad page

	// Slice to store advertisement details
	var ads []Advertisement

	// Create the main context and allocator for the initial page
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false), // Show the browser window
	}

	// Create an allocator and main context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a chromedp context for main page navigation
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Open main page and get list of ad links
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://divar.ir/s/tehran"),
		chromedp.Sleep(2*time.Second),

		// Get all ad titles
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__title")).map(el => el.innerText)`, &adsTitle),
		// Get all ad prices
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__description")).map(el => el.innerText)`, &adsPrice),
		// Get all ad descriptions
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__bottom-description")).map(el => el.innerText)`, &adsDescription),
		// Get all ad links
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__action")).map(el => el.getAttribute('href'))`, &adsLinks),
	)

	if err != nil {
		log.Fatal("Error while scraping main page:", err)
	}

	// Loop through each ad link
	for i := 0; i < len(adsLinks); i++ {
		// Complete the ad link if needed
		adURL := adsLinks[i]
		if !strings.HasPrefix(adURL, "https://") {
			adURL = "https://divar.ir" + adURL
		}

		// Variables to hold ad details
		var details string
		var images []string

		// For each ad, create a separate process to visit and close it independently
		func() {
			// Create a new context for each ad (to open and close each ad in a separate tab)
			adCtx, adCancel := chromedp.NewContext(allocCtx)
			defer adCancel()

			// Open ad page and scrape data
			err = chromedp.Run(adCtx,
				chromedp.Navigate(adURL),
				chromedp.WaitReady(".kt-page-title__description", chromedp.ByQuery), // Wait for the ad description to load

				// Extract detailed description
				chromedp.Text(".kt-page-title__description", &details, chromedp.NodeVisible),

				// Extract image URLs
				chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-image-box__image img")).map(img => img.src)`, &images),
			)
			if err != nil {
				log.Printf("Error visiting ad page %s: %v", adURL, err)
				return
			}

			// Append collected data to ads slice
			ads = append(ads, Advertisement{
				Title:       adsTitle[i],
				Price:       adsPrice[i],
				Description: adsDescription[i],
				Details:     details,
				Images:      images,
			})

			log.Printf("Successfully scraped ad: %s", adsTitle[i])

			// Optional: Add sleep to control request rate
			time.Sleep(1 * time.Second)
		}()
	}

	// Print out all scraped advertisements with additional details
	for _, ad := range ads {
		fmt.Printf("Title: %s\nPrice: %s\nDescription: %s\nDetails: %s\nImages: %v\n\n", ad.Title, ad.Price, ad.Description, ad.Details, ad.Images)
	}
}
