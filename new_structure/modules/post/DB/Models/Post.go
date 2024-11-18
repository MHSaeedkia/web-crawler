package Models

import (
	"project-root/modules/source-site/DB/Models"
	userModels "project-root/modules/user/DB/Models"
	"time"
)

type Post struct {
	ID               int                    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	SrcSitesID       int                    `gorm:"column:source_sites_id;type:bigint(20);NOT NULL"`
	SrcSite          *Models.SourceSite     `gorm:"foreignKey:SrcSitesID;references:ID"`
	UsersID          *int                   `gorm:"column:user_id;type:mediumint;default:null"`
	User             *userModels.User       `gorm:"foreignKey:UsersID;references:ID"`
	Status           int                    `gorm:"column:status;type:tinyint"`
	ExternalSiteID   *string                `gorm:"column:external_site;type:varchar(255);default:null"`
	Title            string                 `gorm:"column:title;type:varchar(255)"`
	Description      *string                `gorm:"column:description;type:varchar(255);default:null"`
	Price            *int64                 `gorm:"column:price;type:bigint;default:null"`
	PriceHistory     map[string]interface{} `gorm:"column:price_history;serializer:json"`
	MainIMG          *string                `gorm:"column:main_img;type:varchar(255);default:null"`
	GalleryIMGs      map[string]interface{} `gorm:"column:gallery_img;serializer:json"`
	SellerName       *string                `gorm:"column:seller_name;type:varchar(45);default:null"`
	LandArea         *float64               `gorm:"column:land_area;type:decimal(10.2);default:null"`
	BuiltYear        *int                   `gorm:"column:built_year;type:int(4);default:null"`
	Rooms            *int                   `gorm:"column:rooms;type:int(4);default:null"`
	IsApartment      *bool                  `gorm:"column:is_apartment;type:tinyint(1);default:null"`
	DealType         *int                   `gorm:"column:deal_type;type:tinyint;default:null"`
	Floors           *int                   `gorm:"column:floors;type:tinyint;default:null"`
	Elevator         *bool                  `gorm:"column:has_elevator;type:tinyint(1);default:null"`
	Storage          *bool                  `gorm:"column:has_storage;type:tinyint(1);default:null"`
	Location         *string                `gorm:"column:location;type:varchar(45);default:null"`
	PostDate         *time.Time             `gorm:"column:post_date;type:datetime;default:null"`
	CityName         *string                `gorm:"column:city_name;type:varchar(100);default:null"`
	NeighborhoodName *string                `gorm:"column:neighborhood_name;type:varchar(100);default:null"`
	IsPublic         bool                   `gorm:"column:is_public;type:tinyint(1);default:1"`
	DeletedAt        *time.Time             `gorm:"column:deleted_at;type:datetime;default:null"`
	CreatedAt        time.Time              `gorm:"column:created_at;type:datetime;NOT NULL"`
	UpdateAt         *time.Time             `gorm:"column:updated_at;type:datetime;NOT NULL"`
}
