package Models

import (
	"time"
)

type CrawlLog struct {
	ID            int        `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	SrcSitesID    int        `gorm:"column:source_sites_id;type:bigint(20);NOT NULL"`
	Status        int        `gorm:"column:status;type:tinyint"`
	TotalRequests int        `gorm:"column:total_requests;type:mediumint"`
	Requests      int        `gorm:"column:sent_requests;type:mediumint"`
	Success       int        `gorm:"column:successful_requests;type:mediumint"`
	Failed        int        `gorm:"column:failed_requests;type:mediumint"`
	RAMUsed       int        `gorm:"column:ram_used;type:mediumint"`
	CPUUsed       int        `gorm:"column:cpu_used;type:mediumint"`
	StartTime     time.Time  `gorm:"column:start_time;type:datetime"`
	EndTime       *time.Time `gorm:"column:end_time;type:datetime;default:null"`
	CreatedAt     time.Time  `gorm:"column:created_at;type:datetime"`
	UpdateAt      time.Time  `gorm:"column:updated_at;type:datetime"`
}
