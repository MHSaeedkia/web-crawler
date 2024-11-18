package connection

import (
	"fmt"
	"log"
	"reflect"
	"time"

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
		&models.Cities{},
		&models.States{},
		&models.Reports{},
		&models.Exports{},
		&models.ReportFilter{},
		&models.SourceSites{},
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

func (c *dbConfig) GenerateFilterQuery(db *gorm.DB, filter models.ReportFilter) *gorm.DB {
	query := db.Model(&models.ReportFilter{})
	if filter.BuiltStart != 0 && filter.BuiltEnd != 0 {
		query = query.Where("built_start >= ? AND built_end <= ?", filter.BuiltStart, filter.BuiltEnd)
	}

	if filter.AreaMin != 0 && filter.AreaMax != 0 {
		query = query.Where("area_min >= ? AND area_max <= ?", filter.AreaMin, filter.AreaMax)
	}

	if filter.PriceMin != 0 && filter.PriceMax != 0 {
		query = query.Where("price_min >= ? AND price_max <= ?", filter.PriceMin, filter.PriceMax)
	}
	val := reflect.ValueOf(filter)
	typ := reflect.TypeOf(filter)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		column := fieldType.Tag.Get("gorm")

		if field.Kind() == reflect.Int && field.Int() == 0 {
			if column == "column:elevator" || column == "column:storage" || column == "column:parking" {
				query = query.Where(fmt.Sprintf("%s = ?", column), field.Int())
			}
			continue
		}
		if isZero(field) {
			continue
		}

		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			query = query.Where(fmt.Sprintf("%s = ?", column), field.Int())
		case reflect.String:
			query = query.Where(fmt.Sprintf("%s LIKE ?", column), "%"+field.String()+"%")
		case reflect.Struct:
			if t, ok := field.Interface().(time.Time); ok && !t.IsZero() {
				query = query.Where(fmt.Sprintf("%s = ?", column), t)
			}
		}
	}

	return query
}
func isZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Ptr, reflect.Interface:
		return value.IsNil()
	case reflect.Struct:
		if t, ok := value.Interface().(time.Time); ok {
			return t.IsZero()
		}
	}
	return false
}


