package Models

import "time"

type User struct {
	ID            int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	IsActive      bool      `gorm:"column:is_active;type:tinyint(1);NOT NULL"`
	Username      string    `gorm:"column:username;type:varchar(255);unique"`
	Password      string    `gorm:"column:password;type:varchar(255)"`
	RoleType      int       `gorm:"column:role_type;type:tinyint;NOT NULL"`
	Email         string    `gorm:"column:email;type:varchar(255)"`
	LastError     time.Time `gorm:"column:last_error;type:datetime"`
	CreatedChatID *int64    `gorm:"column:created_chat_id;type:bigint"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime"`
}
