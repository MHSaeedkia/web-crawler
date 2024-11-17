package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
)

// ConnectDB establishes a connection to the database and returns the DB instance
func ConnectDB() (*gorm.DB, error) {
	endpoint, username, password, port, db := ParseConfig("pkg/config/config.json")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, endpoint, port, db)
	fmt.Println("Connecting with DSN:", dsn)
	// Connecting to the database
	dbConn, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, err
	}

	// Enable logging of SQL queries
	dbConn.LogMode(true)
	return dbConn, nil
}
