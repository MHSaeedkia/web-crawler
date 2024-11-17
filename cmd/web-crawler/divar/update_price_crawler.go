package main

import (
	"fmt"
	"log"

	"github.com/MHSaeedkia/web-crawler/pkg/config"
)

// Post represents the structure of the 'posts' table
type Post struct {
	ID             uint   `gorm:"primaryKey"` // Primary key
	ExternalSiteID string // ExternalSiteID field in the 'posts' table
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
	var externalSiteIDs []string
	for _, post := range posts {
		externalSiteIDs = append(externalSiteIDs, post.ExternalSiteID)
	}

	// Print the list of ExternalSiteID
	fmt.Println("ExternalSiteID List:", externalSiteIDs)
}
