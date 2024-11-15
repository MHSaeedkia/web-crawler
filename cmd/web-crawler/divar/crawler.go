package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/MHSaeedkia/web-crawler/cmd/web-crawler/utils"
	"github.com/MHSaeedkia/web-crawler/internal/models"
	"github.com/MHSaeedkia/web-crawler/pkg/config"
	"github.com/chromedp/chromedp"
	"github.com/jinzhu/gorm"
)

type Site struct {
	BaseURL       string
	LinkSelector  string
	TitleSelector string
	RoomSelector  string
	YearSelector  string
	AreaSelector  string
	PlaceType     string
	ContractType  string
}

type PageData struct {
	Title         string
	TempRoom      string
	Room          int
	TempBuildYear string
	BuildYear     int
	TempArea      string
	Area          int
	TempPrice     string
	Price         int
	PlaceType     string
	ContractType  string
	Elevator      bool
	Parking       bool
	Storage       bool
	Ballcon       bool
	Province      string
	City          string
	ReleaseDate   string
	ImageUrl      string
	Description   string
	TempFloor     string
	Floor         int
	TempRent      string
	Rent          int
	TempDesposit  string
	Desposit      int
}

func main() {
	// Set up a context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Sites configuration
	sites := []Site{
		{
			BaseURL:       "https://divar.ir/s/iran/buy-villa",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
			RoomSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
			YearSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
			AreaSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
			PlaceType:     "villa",
			ContractType:  "buy",
		},
		{
			BaseURL:       "https://divar.ir/s/iran/buy-apartment",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
			RoomSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
			YearSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
			AreaSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
			PlaceType:     "apartment",
			ContractType:  "buy",
		},
		{
			BaseURL:       "https://divar.ir/s/iran/rent-villa",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
			RoomSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
			YearSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
			AreaSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
			PlaceType:     "villa",
			ContractType:  "rent",
		},
		{
			BaseURL:       "https://divar.ir/s/iran/rent-apartment",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
			RoomSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
			YearSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
			AreaSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
			PlaceType:     "apartment",
			ContractType:  "rent",
		},

		// Add other sites as needed
	}

	var wg sync.WaitGroup

	for _, site := range sites {
		contractType, placeType := utils.ExtractContractAndPlaceType(site.BaseURL)

		wg.Add(1)
		go func(site Site, contractType, placeType string) {
			defer wg.Done()
			scrapeSite(ctx, site, contractType, placeType, db)
		}(site, contractType, placeType)
	}

	// Wait for all scraping routines to complete
	go func() {
		wg.Wait()
		stop() // Signal that all work is done
	}()

	// Block until we receive an interrupt or all work is completed
	<-ctx.Done()
	log.Println("Received shutdown signal, waiting for all tasks to complete...")

	// Wait for any remaining work
	wg.Wait()
	log.Println("All tasks completed. Program shutting down gracefully.")
}

func scrapeSite(ctx context.Context, site Site, contractType, placeType string, db *gorm.DB) {
	siteCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	log.Printf("Navigating to base URL: %s", site.BaseURL)
	err := chromedp.Run(siteCtx,
		chromedp.Navigate(site.BaseURL),
		chromedp.WaitVisible(site.LinkSelector, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Printf("Failed to load base URL %s: %v", site.BaseURL, err)
		return
	}
	log.Printf("Successfully loaded base URL: %s", site.BaseURL)

	var links []string
	err = chromedp.Run(siteCtx,
		chromedp.Evaluate(fmt.Sprintf(`Array.from(document.querySelectorAll("%s")).map(a => a.href)`, site.LinkSelector), &links),
	)
	if err != nil {
		log.Printf("Failed to retrieve links from %s: %v", site.BaseURL, err)
		return
	}
	for i, link := range links {
		select {
		case <-ctx.Done():
			log.Println("Shutdown signal received, stopping further processing")
			return
		default:
			data, err := scrapeLink(siteCtx, link, site, contractType, placeType, db)
			if err != nil {
				log.Printf("Failed to scrape page %s: %v", link, err)
				continue
			}
			log.Printf("Extracted data from page %d: %+v", i+1, data)
		}
	}
	log.Printf("Completed scraping all links for site: %s", site.BaseURL)
}

func scrapeLink(ctx context.Context, link string, site Site, contractType, placeType string, db *gorm.DB) (PageData, error) {
	var data PageData
	var result string
	data.ContractType = contractType
	data.PlaceType = placeType
	time.Sleep(10 * time.Second)
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	var elements []string
	var contentSlice []string

	err1 := chromedp.Run(timeoutCtx,
		chromedp.Navigate(link),
		chromedp.WaitVisible(site.TitleSelector, chromedp.ByQuery),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("p.kt-unexpandable-row__value")).map(e => e.textContent.trim())`, &elements),
		chromedp.Evaluate(`Array.from(document.querySelectorAll(".kt-body--stable")).map(el => el.textContent)`, &contentSlice),
		chromedp.Text(site.TitleSelector, &data.Title, chromedp.ByQuery),
		chromedp.Text(site.RoomSelector, &data.TempRoom, chromedp.ByQuery),
		chromedp.Text(site.YearSelector, &data.TempBuildYear, chromedp.ByQuery),
		chromedp.Text(site.AreaSelector, &data.TempArea, chromedp.ByQuery),
		chromedp.Text(`div.kt-page-title__subtitle.kt-page-title__subtitle--responsive-sized`, &result, chromedp.ByQuery),
		chromedp.EvaluateAsDevTools(`
			(function() {
				let img = document.querySelector("#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-6.kt-offset-1 > section:nth-child(1) > div > div > div.keen-slider.kt-base-carousel__slides.slides-d6304 > div > figure > div > picture > img");
				return img ? img.src : "";
			})()`, &data.ImageUrl),
		chromedp.Text(`p.kt-description-row__text.kt-description-row__text--primary`, &data.Description, chromedp.ByQuery),
	)
	if err1 != nil {
		return PageData{}, err1
	}
	parts := strings.Split(result, " در ")
	data.ReleaseDate = parts[0]
	Location := strings.Split(parts[1], "،")
	if len(Location) > 1 {
		data.Province = Location[0]
		data.City = Location[1]
	} else {
		data.Province = ""
		data.City = parts[1]
	}

	// Convert TempRoom, TempBuildYear, TempArea, and TempPrice
	linkParts := strings.Split(link, "/")
	linkID := linkParts[len(linkParts)-1]
	data.Room = utils.ConvertToInt(data.TempRoom)
	data.BuildYear = utils.ConvertToInt(data.TempBuildYear)
	data.Area = utils.ConvertToInt(data.TempArea)
	if placeType == "villa" {
		if contractType == "buy" {
			if len(elements) == 3 {
				data.TempPrice = elements[1]
			} else if len(elements) > 3 {
				data.TempPrice = elements[2]
			}
			data.Price = utils.ConvertToInt(data.TempPrice)
			data.Parking = utils.ConvertFeatureToBool(contentSlice[0])
			data.Storage = utils.ConvertFeatureToBool(contentSlice[1])
			data.Ballcon = utils.ConvertFeatureToBool(contentSlice[2])
			data.Elevator = false
			data.Floor = 0
			data.Rent = 0
			data.Desposit = 0
		} else {
			if len(elements) == 4 {
				data.TempDesposit = elements[1]
				data.TempRent = elements[2]
			} else {
				return data, errors.New("this is an error message")
			}
			data.Price = 0
			data.Parking = utils.ConvertFeatureToBool(contentSlice[0])
			data.Storage = utils.ConvertFeatureToBool(contentSlice[1])
			data.Ballcon = utils.ConvertFeatureToBool(contentSlice[2])
			data.Elevator = false
			data.Floor = 0
			data.Rent = utils.ConvertToInt(data.TempRent)
			data.Desposit = utils.ConvertToInt(data.TempDesposit)
		}
	} else {
		if contractType == "buy" {
			data.TempPrice = elements[0]
			data.TempFloor = elements[2]
			data.Price = utils.ConvertToInt(data.TempPrice)
			data.Elevator = utils.ConvertFeatureToBool(contentSlice[0])
			data.Parking = utils.ConvertFeatureToBool(contentSlice[1])
			data.Storage = utils.ConvertFeatureToBool(contentSlice[2])
			data.Ballcon = false
			data.Floor = utils.ConvertFloor(data.TempFloor)
		} else {
			if len(elements) == 4 {
				data.TempDesposit = elements[0]
				data.TempRent = elements[1]
			} else {
				return data, errors.New("this is an error message")
			}
			data.TempFloor = elements[3]
			data.Price = 0
			data.Elevator = utils.ConvertFeatureToBool(contentSlice[0])
			data.Parking = utils.ConvertFeatureToBool(contentSlice[1])
			data.Storage = utils.ConvertFeatureToBool(contentSlice[2])
			data.Ballcon = false
			data.Floor = utils.ConvertFloor(data.TempFloor)
			data.Rent = utils.ConvertToInt(data.TempRent)
			data.Desposit = utils.ConvertToInt(data.TempDesposit)
		}

	}
	post := models.Posts{
		SourceSiteId:   1,
		ExternalSiteID: linkID,
		Title:          data.Title,
		Description:    data.Description,
		Price:          data.Price,
		PriceHistory:   "",
		MainIMG:        data.ImageUrl,
		GalleryIMG:     "",
		SellerName:     "Unknown",
		LandArea:       float64(data.Area),
		BuiltYear:      data.BuildYear,
		RoomCount:      data.Room,
		IsApartment:    false,
		DealType:       1,
		Floors:         data.Floor,
		Elevator:       data.Elevator,
		Storage:        data.Storage,
		//Ballcon:	false,
		//Parking:  false,
		Location:         "Sample location",
		PostDate:         time.Now(),
		CityName:         data.Province,
		NeighborhoodName: data.City,
	}

	err1 = db.Create(&post).Error
	if err1 != nil {
		log.Printf("Error saving post: %v", err1)
		return PageData{}, err1
	}

	log.Printf("Saved post with ID %d to database", post.ID)

	return data, nil
}
