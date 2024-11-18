package Connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project-root/sys-modules/env"
	"sync"
)

func NewMysqlDbConnection() *gorm.DB {
	var (
		once          sync.Once
		connection_db *gorm.DB
	)
	once.Do(func() {
		dsn := getConnectionString()
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}
		connection_db = db
	})
	return connection_db
}

func getConnectionString() string {
	username := env.Env("DB_USERNAME")
	password := env.Env("DB_PASSWORD")
	host := env.Env("DB_HOST")
	port := env.Env("DB_PORT")
	dbname := env.Env("DB_NAME")
	charset := "utf8mb4"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, dbname, charset)
}
