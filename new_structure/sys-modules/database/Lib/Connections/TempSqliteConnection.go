package Connections

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

func NewInMemorySqliteDbConnection() *gorm.DB {
	var (
		once          sync.Once
		connection_db *gorm.DB
	)
	once.Do(func() {
		// ایجاد اتصال به دیتابیس حافظه‌ای (در حافظه RAM)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to in-memory sqlite database: " + err.Error())
		}
		connection_db = db
	})
	return connection_db
}
