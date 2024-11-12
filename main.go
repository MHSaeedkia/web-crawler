package main

import (
	"log"

	"github.com/MHSaeedkia/web-crawler/cmd/connection"
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
}
