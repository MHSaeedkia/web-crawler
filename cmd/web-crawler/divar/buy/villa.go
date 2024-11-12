package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
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
	//ParkingSelector string
	//CallerSelector  string
	//BallconSelector string
	//Attribiute      string
}

type PageData struct {
	Title        string
	Room         string
	BuildYear    string
	Area         string
	Price        string
	PlaceType    string
	ContractType string
	Parking      string
	Cellar       string
	Ballcon      string
}

func main() {
	// Set up a context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Sites configuration
	sites := []Site{
		{
			BaseURL:       "https://divar.ir/s/iran/buy-villa",
			LinkSelector:  "a.kt-post-card__action",
			TitleSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.kt-page-title > div > h1",
			RoomSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(3)",
			YearSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(2)",
			AreaSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(1) > tbody > tr > td:nth-child(1)",
			//BallconSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(3)",
			//ParkingSelector: "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(1)",
			//CallerSelector:  "#app > div.container--has-footer-d86a9.kt-container > div > main > article > div > div.kt-col-5 > section:nth-child(1) > div.post-page__section--padded > table:nth-child(10) > tbody > tr > td:nth-child(2)",
			//Attribiute:      `//*[contains(concat(" ", @class, " "), concat(" ", "kt-body--stable", " "))]`,
			PlaceType:    "villa",
			ContractType: "buy",
		},
		// Add other sites as needed
	}

	var wg sync.WaitGroup

	for _, site := range sites {
		contractType, placeType := extractContractAndPlaceType(site.BaseURL)

		wg.Add(1)
		go func(site Site, contractType, placeType string) {
			defer wg.Done()
			scrapeSite(ctx, site, contractType, placeType)
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

func extractContractAndPlaceType(url string) (contractType, placeType string) {
	parts := strings.Split(url, "/")
	if len(parts) >= 6 {
		types := strings.Split(parts[5], "-")
		if len(types) >= 2 {
			contractType = types[0]
			placeType = types[1]
		}
	}
	return
}

func scrapeSite(ctx context.Context, site Site, contractType, placeType string) {
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
			data, err := scrapeLink(siteCtx, link, site, contractType, placeType)
			if err != nil {
				log.Printf("Failed to scrape page %s: %v", link, err)
				continue
			}
			log.Printf("Extracted data from page %d: %+v", i+1, data)
		}
	}
	log.Printf("Completed scraping all links for site: %s", site.BaseURL)
}

func scrapeLink(ctx context.Context, link string, site Site, contractType, placeType string) (PageData, error) {
	var data PageData
	var divCount int
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
	)
	if err1 != nil {
		fmt.Println(divCount)
		fmt.Println(data)
		return PageData{}, err1
	} else {
		err := chromedp.Run(timeoutCtx,
			chromedp.Navigate(link),
			chromedp.WaitVisible(site.TitleSelector, chromedp.ByQuery),
			chromedp.Text(site.TitleSelector, &data.Title, chromedp.ByQuery),
			chromedp.Text(site.RoomSelector, &data.Room, chromedp.ByQuery),
			chromedp.Text(site.YearSelector, &data.BuildYear, chromedp.ByQuery),
			chromedp.Text(site.AreaSelector, &data.Area, chromedp.ByQuery),
		)
		if err != nil {
			return PageData{}, err
		}

		if len(elements) == 3 {
			data.Price = elements[1]
		} else if len(elements) > 3 {
			data.Price = elements[2]
		}
		data.Parking = contentSlice[0]
		data.Cellar = contentSlice[1]
		data.Ballcon = contentSlice[2]
	}
	return data, nil
}
