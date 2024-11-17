package models

import (
	"time"
)

type Posts struct {
	ID             *int                    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	SRCSiteID      int                     `gorm:"column:user_id"`
	SRCSite        SourceSites             `gorm:"foreignKey:SRCSiteID;references:ID"`
	UserID         int                     `gorm:"column:user_id"`
	User           Users                   `gorm:"foreignKey:UserID;references:ID"`
	Status         *int                    `gorm:"column:name;type:tinyint"`
	ExternalSiteID *string                 `gorm:"column:external_site;type:varchar(255)"`
	Title          *string                 `gorm:"column:title;type:varchar(255)"`
	Description    *string                 `gorm:"column:description;type:text"`
	Price          *int                    `gorm:"column:project;type:bigint"`
	PriceHistory   *map[string]interface{} `gorm:"column:price_history;serializer:json"`
	MainIMG        *string                 `gorm:"column:main_img;type:varchar(255)"`
	GalleryIMGs    *map[string]interface{} `gorm:"column:gallery_img;serializer:json"`
	SellerName     *string                 `gorm:"column:seller_name;type:varchar(45)"`
	LandArea       *float64                `gorm:"column:land_area;type:decimal(10.2)"`
	BuiltYear      *int                    `gorm:"column:built_year;type:int(4)"`
	Rooms          *int                    `gorm:"column:rooms;type:tinyint"`
	IsApartment    *int                    `gorm:"column:is_apartment;type:tinyint(1)"`
	DealType       *int                    `gorm:"column:deal_type;type:tinyint"`
	Floors         *int                    `gorm:"column:floors;type:tinyint"`
	Elevator       *int                    `gorm:"column:has_elevator;type:tinyint(1)"`
	Storage        *int                    `gorm:"column:has_storage;type:tinyint(1)"`
	Location       *string                 `gorm:"column:location;type:varchar(45)"`
	City           *string                 `gorm:"column:city;type:varchar(255)"`
	Neighbourhood  *string                 `gorm:"column:neighbourhood;type:varchar(255)"`
	PostDate       *time.Time              `gorm:"column:post_date;type:datetime"`
	DeletedAt      *time.Time              `gorm:"column:deleted_at;type:datetime"`
	CreatedAt      *time.Time              `gorm:"column:created_at;type:datetime"`
	UpdatedAt      *time.Time              `gorm:"column:updated_at;type:datetime"`
}
