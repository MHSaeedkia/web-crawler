package Models

import (
	"time"
)

type ReportFilter struct {
	ID             int        `gorm:"column:id;primaryKey;AUTOINCREMENT"`
	ReportID       int        `gorm:"column:reports_id;type:bigint"`
	Report         *Report    `gorm:"foreignKey:ReportID;references:ID"`
	BuiltStart     *int       `gorm:"column:built_start;type:smallint;default:null"`
	BuiltEnd       *int       `gorm:"column:built_end;type:smallint;default:null"`
	AreaMin        *int       `gorm:"column:area_min;type:smallint;default:null"`
	AreaMax        *int       `gorm:"column:area_max;type:smallint;default:null"`
	PriceMin       *int       `gorm:"column:price_min;type:bigint;default:null"`
	PriceMax       *int       `gorm:"column:price_max;type:bigint;default:null"`
	RoomCountMin   *int       `gorm:"column:room_count_min;type:tinyint;default:null"`
	RoomCountMax   *int       `gorm:"column:room_count_max;type:tinyint;default:null"`
	DealType       *int       `gorm:"column:deal_type;type:tinyint;default:null"`
	FloorCountMin  *int       `gorm:"column:floors_min;tinyint;default:null"`
	FloorCountMax  *int       `gorm:"column:floors_max;tinyint;default:null"`
	Elevator       *int       `gorm:"column:elevator;type:tinyint;default:null"`
	Storage        *int       `gorm:"column:storage;type:tinyint;default:null"`
	Parking        *int       `gorm:"column:parking;type:tinyint;default:null"`
	Location       *string    `gorm:"column:location;type:varchar(45);default:null"`
	LocationRadius *int       `gorm:"column:location_radius;type:bigint;default:null"`
	IsApartment    *int       `gorm:"column:is_apartment;type:tinyint;default:null"`
	PostStartDate  *time.Time `gorm:"column:post_start;type:datetime;default:null"`
	PostEndDate    *time.Time `gorm:"column:post_end;type:datetime;default:null"`
}
