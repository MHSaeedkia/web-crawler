package models

import (
	"time"
)

type UserBookMarks struct {
	ID            int       `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	User_id    int       `gorm:"column:source_site_id;type:bigint(20);NOT NULL"`
	Status        int       `gorm:"column:name;type:tinyint"`
	TotalRequests int       `gorm:"column:total_requests;type:mediumint"`
	Requests      int       `gorm:"column:sent_requests;type:mediumint"`
	Success       int       `gorm:"column:Successful_requests;type:mediumint"`
	Faild         int       `gorm:"column:faild_requests;type:mediumint"`
	RAMUsed       int       `gorm:"column:ram_used;type:mediumint"`
	CPUUsed       int       `gorm:"column:cpu_used;type:decimal(10,2)"`
	StartTime     time.Time `gorm:"column:start_time;type:datetime"`
	EndTime       time.Time `gorm:"column:end_time;type:datetime"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime"`
	UpdateAt      time.Time `gorm:"column:updated_at;type:datetime"`
}