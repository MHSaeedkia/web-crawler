package models

import (
	"time"
)

type ReportFilter struct {
	ID             int       `gorm:"column:id;primaryKey;AUTOINCREMENT"`
	ReportID       int       `gorm:"column:reports_id;type:bigint"`
	CitiesID       int       `gorm:"column:cities_id;type:bigint"`
	BuiltStart     int       `gorm:"column:built_start;type:smallint"`
	BuiltEnd       int       `gorm:"column:built_end;type:smallint"`
	AreaMin        int       `gorm:"column:area_min;type:smallint"`
	AreaMax        int       `gorm:"column:area_max;type:smallint"`
	PriceMin       int       `gorm:"column:price_min;type:bigint"`
	PriceMax       int       `gorm:"column:price_max;type:bigint"`
	RoomCountMin   int       `gorm:"column:room_count_min;type:tinyint"`
	RoomCountMax   int       `gorm:"column:room_count_max;type:tinyint"`
	DealType       int       `gorm:"column:deal_type;type:tinyint"`
	FloorCountMin  int       `gorm:"column:floors_min;tinyint"`
	FloorCountMax  int       `gorm:"column:floors_max;tinyint"`
	Elevator       int       `gorm:"column:elevator;type:tinyint"`
	Storage        int       `gorm:"column:storage;type:tinyint"`
	Parking        int       `gorm:"column:parking;type:tinyint"`
	Location       string    `gorm:"column:location;type:varchar(45)"`
	LocationRadius int       `gorm:"column:location_radius;type:bigint"`
	IsApartment    int       `gorm:"column:is_apartment;type:tinyint"`
	PostStartDate  time.Time `gorm:"column:post_start;type:datetime"`
	PostEndDate    time.Time `gorm:"column:post_end;type:datetime"`
}
