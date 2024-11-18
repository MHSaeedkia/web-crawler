package models

import "time"

type Users struct {
	ID            int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	IsActive      int       `gorm:"column:is_active;type:tinyint;NOT NULL"`
	Username      string    `gorm:"column:username;type:varchar(255)"`
	Password      string    `gorm:"column:password;type:varchar(255)"`
	ChatID        int       `gorm:"column:chat_id;type:varchar(255)"`
	RoleType      int       `gorm:"column:role_type;type:tinyint;NOT NULL"`
	Email         string    `gorm:"column:email;type:varchar(255)"`
	LastError     time.Time `gorm:"column:last_error;type:datetime"`
	CreatedChatID time.Time `gorm:"column:created_chat_id;type:varchar(45)"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime"`
}
