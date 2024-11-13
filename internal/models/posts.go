package models

import (
	"time"
)

type Posts struct {
	ID             int     `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	SourceSiteId   int     `gorm:"column:source_site_id;type:bigint(20);NOT NULL"`
	CitiesID       int     `gorm:"column:city_id;type:mediumint;NOT NULL"`
	UsersID        int     `gorm:"column:user_id;type:mediumint"`
	Status         int     `gorm:"column:name;type:tinyint"`
	ExternalSiteID string  `gorm:"column:external_site_id;type:varchar(255)"`
	Title          string  `gorm:"column:title;type:varchar(45)"`
	Description    string  `gorm:"column:description;type:varchar(255)"`
	Price          int     `gorm:"column:project;type:bigint"`
	PriceHistory   string  `json:"price_history"`
	MainIMG        string  `gorm:"column:main_img;type:varchar(255)"`
	GalleryIMG     string  `json:"gallery_img"`
	SellerName     string  `gorm:"column:seller_name;type:varchar(45)"`
	LandArea       float64 `gorm:"column:land_area;type:decimal(10.2)"`
	BuiltYear      int     `gorm:"column:built_year;type:int(4)"`
	// Rooms          int       `gorm:"column:built_year;type:int(4)"`
	IsApartment bool      `gorm:"column:is_apartment;type:tinyint(1)"`
	DealType    int       `gorm:"column:deal_type;type:tinyint"`
	Floors      int       `gorm:"column:floors;type:tinyint"`
	Elevator    bool      `gorm:"column:has_elevator;type:tinyint(1)"`
	Storage     bool      `gorm:"column:has_storage;type:tinyint(1)"`
	Location    string    `gorm:"column:location;type:varchar(45)"`
	PostDate    time.Time `gorm:"column:post_date;type:datetime"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:datetime"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime"`
	UpdateAt    time.Time `gorm:"column:updated_at;type:datetime"`
}
