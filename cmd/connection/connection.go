package connection

import (
	"fmt"
	"log"

	"github.com/MHSaeedkia/web-crawler/input"
	"github.com/MHSaeedkia/web-crawler/internal/models"
	"github.com/MHSaeedkia/web-crawler/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseConnection struct {
	*gorm.DB
}

type dbConfig struct {
	address  string
	user     string
	password string
	port     int
	name     string
}

func GetDBConfig() *dbConfig {
	configPath := input.GetInput()
	ifExists := config.CheckConfigFile(configPath)
	if ifExists {
		dbAddress, dbUser, dbPassword, dbPort, dbName := config.ParseConfig(configPath)
		return &dbConfig{address: dbAddress, user: dbUser, password: dbPassword, port: dbPort, name: dbName}
	}
	log.Println("Config file not found. Using default values.")
	return &dbConfig{}
}
func (c *dbConfig) AutoMigrate(dbConn *databaseConnection) error {
	tables := []interface{}{
		&models.CrawlLogs{},
		&models.Posts{},
		&models.TelSessions{},
		&models.UserBookMarks{},
		&models.Users{},
	}

	for _, table := range tables {
		err := dbConn.AutoMigrate(table)
		if err != nil {
			log.Fatalf("Failed to migrate database for table %v: %v", table, err)
			return err
		}
	}
	fmt.Println("Database migration completed for all tables")
	return nil
}

func (c *dbConfig) Connect() (*databaseConnection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.user,
		c.password,
		c.address,
		c.port,
		c.name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &databaseConnection{DB: db}, nil
}
