package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MHSaeedkia/web-crawler/pkg/config"
	"github.com/MHSaeedkia/web-crawler/cmd/web-crawler/utils"
	"github.com/chromedp/chromedp"
	"gorm.io/gorm"
)

// Post represents the structure of the 'posts' table
type Post struct {
	ID           uint   `gorm:"primaryKey"` // Primary key
	ExternalSiteID string
	Price        int
	PriceHistory string
}

// ScrapeAndCheckPrice function as previously implemented
func ScrapeAndCheckPrice(link string, givenPrice int, placeType bool, contractType int) (int, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	titleSelector := "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1"
	var elements []string
	var scrapedPriceText string
	var price int

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(titleSelector, chromedp.ByQuery),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("p.kt-unexpandable-row__value")).map(e => e.textContent.trim())`, &elements),
	)
	if err != nil {
		fmt.Println("Failed to scrape:", err)
		return 0, fmt.Errorf("failed to scrape the link: %w", err)
	}

	if !placeType { // Villa
		if contractType == 1 { // Buy
			if len(elements) == 3 {
				TempPrice = elements[1]
			} else if len(elements) > 3 {
				TempPrice = elements[2]
			}
			Price = utils.ConvertToInt(TempPrice)
		} else { // Rent
			if len(elements) == 4 {
				TempDesposit = elements[1]
			} else {
				return 0, fmt.Errorf("invalid data structure")
			}
			Price = utils.ConvertToInt(TempDesposit)
		}
	} else { // Apartment
		if contractType == 1 { // Buy
			TempPrice = elements[0]
			Price = utils.ConvertToInt(TempPrice)
		} else { // Rent
			if len(elements) == 4 {
				TempDesposit = elements[0]
			} else {
				return 0, fmt.Errorf("invalid data structure")
			}
			Price = utils.ConvertToInt(TempDesposit)
		}
	}

	if Price == givenPrice {
		return 0, nil
	}
	return Price, nil
}


func updatePriceCrawler() {
	// Connect to the database
	db, err := config.ConnectDB() // Assumes config.ConnectDB() is implemented as in previous examples
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Fetch all ExternalSiteID values from the posts table
	var posts []Post
	if err := db.Select("external_site_id").Find(&posts).Error; err != nil {
		log.Fatal("Failed to fetch ExternalSiteID from posts:", err)
		return
	}

	// Extract ExternalSiteID into a list
	for _, post := range posts {
		link := post.ExternalSiteID

		newPrice, err := ScrapeAndCheckPrice(link, post.Price, post.IsApartment, post.DealType) // Example: Assume placeType=true, contractType=1
		if err != nil {
			log.Printf("Failed to scrape ExternalSiteID %s: %v\n", post.ExternalSiteID, err)
			continue
		}
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