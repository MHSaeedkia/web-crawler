package models

import (
	"time"
)

type UserBookMarks struct {
	ID           int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	UserID       *Users    `gorm:"column:User_id;type:mediumint;NOT NULL;foreignKey:UserID"`
	PostID       *Posts    `gorm:"column:post_id;type:mediumint;NOT NULL;foreignKey:PostID"`
	BookMarkedAt time.Time `gorm:"column:bookmarked_at;type:datetime"`
	IsPublic     int       `gorm:"column:is_public;type:tinyint"`
}
