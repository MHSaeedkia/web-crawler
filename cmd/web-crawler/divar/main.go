package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
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

// Function to scrape details of a single ad
func scrapeAd(adURL string, title, price, description string, adCh chan<- Advertisement, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done

	// Define a struct to store ad details
	var ad Advertisement

	// Set up a timeout context for the ad page (e.g., 10 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure this context gets canceled after completion

	// Create a new allocator with visible window for each ad to avoid issues
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
	}
	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	// Create a new browser context for each ad page
	adCtx, adCancel := chromedp.NewContext(allocCtx)
	defer adCancel() // Cancel ad context to free up resources

	// Initialize variables for ad details
	var details string
	var images []string

	// Run scraping tasks for this ad
	err := chromedp.Run(adCtx,
		chromedp.Navigate(adURL),
		chromedp.Sleep(2*time.Second), // Allow some time for page load
		chromedp.Text(".kt-page-title__description", &details, chromedp.NodeVisible),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-image-box__image img")).map(img => img.src)`, &images),
	)

	if err != nil {
		log.Printf("Error visiting ad page %s: %v", adURL, err)
		return
	}

	// Assign scraped data to ad struct
	ad = Advertisement{
		Title:       title,
		Price:       price,
		Description: description,
		Details:     details,
		Images:      images,
	}

	// Send ad to channel
	adCh <- ad
}

func main() {
	// Main context and allocator for the main page
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
	}
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create a new browser context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Define slices to store scraped data from the main page
	var adsTitle []string
	var adsPrice []string
	var adsDescription []string
	var adsLinks []string

	// Navigate to the main page and collect links
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://divar.ir/s/tehran"),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__title")).map(el => el.innerText)`, &adsTitle),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__description")).map(el => el.innerText)`, &adsPrice),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__bottom-description")).map(el => el.innerText)`, &adsDescription),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-post-card__action")).map(el => el.getAttribute('href'))`, &adsLinks),
	)

	if err != nil {
		log.Fatal("Error while scraping main page:", err)
	}

	// Channel to collect ads from goroutines
	adCh := make(chan Advertisement, len(adsLinks))
	var wg sync.WaitGroup

	// Limit the number of concurrent goroutines to avoid opening too many tabs
	semaphore := make(chan struct{}, 5) // Limit to 5 concurrent tabs

	for i := 0; i < len(adsTitle); i++ {
		adURL := adsLinks[i]
		if !strings.HasPrefix(adURL, "https://") {
			adURL = "https://divar.ir" + adURL
		}

		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot

		go func(i int, adURL string) {
			defer func() { <-semaphore }() // Release the slot when done

			scrapeAd(adURL, adsTitle[i], adsPrice[i], adsDescription[i], adCh, &wg)
		}(i, adURL)
	}

	// Wait for all scraping goroutines to finish
	go func() {
		wg.Wait()
		close(adCh) // Close the channel after all goroutines are done
	}()

	// Print out all scraped advertisements with additional details
	for ad := range adCh {
		fmt.Printf("Title: %s\nPrice: %s\nDescription: %s\nDetails: %s\nImages: %v\n\n", ad.Title, ad.Price, ad.Description, ad.Details, ad.Images)
	}
}
