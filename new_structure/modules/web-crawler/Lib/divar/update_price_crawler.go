package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MHSaeedkia/web-crawler/cmd/web-crawler/utils"
	"github.com/MHSaeedkia/web-crawler/pkg/config"
	"github.com/chromedp/chromedp"
)

// Post represents the structure of the 'posts' table
type Post struct {
	ID             uint `gorm:"primaryKey"` // Primary key
	ExternalSiteID string
	Price          int
	PriceHistory   string
	IsApartment    bool
	DealType       int
}

// ScrapeAndCheckPrice function as previously implemented
func ScrapeAndCheckPrice(link string, givenPrice int, placeType bool, contractType int) (int, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var elements []string
	var tempPrice string
	var price int

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(`#app > div.container--has-footer-d86a9.kt-container`, chromedp.ByQuery),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("p.kt-unexpandable-row__value")).map(e => e.textContent.trim())`, &elements),
	)
	if err != nil {
		fmt.Println("Failed to scrape:", err)
		return 0, fmt.Errorf("failed to scrape the link: %w", err)
	}

	// Extract price logic here based on `placeType` and `contractType`
	if placeType { //apartement
		tempPrice = elements[0]
		price = utils.ConvertToInt(tempPrice)
	} else {
		// Villa logic
		if contractType == 1 {
			if len(elements) == 3 {
				tempPrice = elements[1]
			} else if len(elements) > 3 {
				tempPrice = elements[2]
			}
			price = utils.ConvertToInt(tempPrice)
		} else {
			if len(elements) == 4 {
				tempPrice = elements[0]
				price = utils.ConvertToInt(tempPrice)
			} else {
				return 0, nil
			}
		}

	}

	if price == givenPrice {
		return 0, nil // No change
	}
	return price, nil // Return new price if changed
}

func updatePriceCrawler() {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Fetch all posts with ExternalSiteID and Price
	var posts []Post
	if err := db.Select("id, external_site_id, price, price_history").Find(&posts).Error; err != nil {
		log.Fatal("Failed to fetch posts:", err)
		return
	}

	for _, post := range posts {
		// Construct the link from the ExternalSiteID
		link := fmt.Sprintf("https://divar.ir/v/%s", post.ExternalSiteID)
		placetype := post.IsApartment
		contractType := post.DealType

		// Scrape the current price
		newPrice, err := ScrapeAndCheckPrice(link, post.Price, placetype, contractType) // Example: Assume placeType=true, contractType=1
		if err != nil {
			log.Printf("Failed to scrape ExternalSiteID %s: %v\n", post.ExternalSiteID, err)
			continue
		}

		// If the price has changed, update the database
		if newPrice != 0 {
			// Update price history
			newPriceHistory := fmt.Sprintf("%s;%d:%s", post.PriceHistory, post.Price, time.Now().Format("2006-01-02 15:04:05"))

			// Update record
			if err := db.Model(&post).Updates(map[string]interface{}{
				"price":         newPrice,
				"price_history": newPriceHistory,
			}).Error; err != nil {
				log.Printf("Failed to update price for ExternalSiteID %s: %v\n", post.ExternalSiteID, err)
			} else {
				log.Printf("Price updated for ExternalSiteID %s: Old Price: %d, New Price: %d\n", post.ExternalSiteID, post.Price, newPrice)
			}
		}
	}
}

func main() {
	updatePriceCrawler()
}
