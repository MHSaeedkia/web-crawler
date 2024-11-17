package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MHSaeedkia/web-crawler/cmd/connection"
	"github.com/MHSaeedkia/web-crawler/internal/models"
)

func main() {
	dbConfig := connection.GetDBConfig()
	dbConn, err := dbConfig.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	migrationerr := dbConfig.AutoMigrate(dbConn)
	if migrationerr != nil {
		log.Fatalf("Couldn't migrate tables to database: %v", migrationerr)
	}
	filter := models.ReportFilter{
		BuiltStart:    2000,
		BuiltEnd:      2020,
		AreaMin:       50,
		AreaMax:       200,
		Location:      "Downtown",
		Elevator:      1, // '0' as a valid search value for specific fields like Elevator, Storage, Parking
		PostStartDate: time.Now().AddDate(-1, 0, 0),
	}

	// Generate and print the query
	query := dbConfig.GenerateFilterQuery(dbConn.DB, filter)
	fmt.Println(query.Statement.Clauses)

}
