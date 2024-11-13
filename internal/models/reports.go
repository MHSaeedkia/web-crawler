package models

import "time"

type Reports struct {
	ID        int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	UserID    string    `gorm:"column:user_id;type:varchar(45)"`
	Title     int    `gorm:"column:is_active;type:tinyint"`
	IsActive  int    `gorm:"column:is_notification;type:tinyint"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
}
