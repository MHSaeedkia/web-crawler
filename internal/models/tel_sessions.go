package models

type TelSessions struct {
	ID           int                    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	LogingUserID int                    `gorm:"column:logined_users_id;type:mediumint;NOT NULL"`
	ChatID       string                 `gorm:"column:chat_id;type:varchar(100)"`
	CurrentPage  int                    `gorm:"column:current_page_num;type:mediumint"`
	Temp_data    map[string]interface{} `gorm:"column:temp_data;serializer:json"`
}
