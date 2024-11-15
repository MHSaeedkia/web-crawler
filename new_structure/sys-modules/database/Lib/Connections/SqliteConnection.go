package Connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

func NewSqliteDbConnection() *gorm.DB {
	var (
		once          sync.Once
		connection_db *gorm.DB
	)
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to sqlite database: " + err.Error())
		}
		connection_db = db
	})
	return connection_db
}
