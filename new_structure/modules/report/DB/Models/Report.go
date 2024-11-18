package Models

import (
	"project-root/modules/user/DB/Models"
	"time"
)

type Report struct {
	ID             int          `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	UserID         *int         `gorm:"column:users_id;type:mediumint"`
	User           *Models.User `gorm:"foreignKey:UserID;references:ID"`
	Title          string       `gorm:"column:title;type:varchar(255);"`
	IsNotification int          `gorm:"column:is_notification;type:tinyint"`
	CreatedAt      time.Time    `gorm:"column:created_at;type:datetime"`
	DeletedAt      *time.Time   `gorm:"column:deleted_at;type:datetime;default:null"`
}
