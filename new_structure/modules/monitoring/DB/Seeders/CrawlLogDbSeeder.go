package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/monitoring/DB/Models"
	"project-root/modules/monitoring/Enums"
	"project-root/modules/monitoring/Facades"
	"project-root/sys-modules/database/Lib"
	"time"
)

type CrawlLogDbSeeder struct{}

func (s CrawlLogDbSeeder) Name() string {
	return "400_crawl_logs_table"
}

func (s CrawlLogDbSeeder) Handle(db *gorm.DB) {
	Facades.CrawlLogRepo().Truncate()
	now := time.Now()
	endTime := now.Add(-92 * time.Hour)
	items := []Models.CrawlLog{
		{
			SrcSitesID:    1,
			Status:        Enums.SuccessfulCrawlLogStatus,
			TotalRequests: 1500,
			Requests:      1500,
			Success:       1000,
			Failed:        500,
			RAMUsed:       32600,
			CPUUsed:       4632,
			StartTime:     now.Add(-720 * time.Hour),
			EndTime:       nil,
			CreatedAt:     now.Add(-92 * time.Hour),
			UpdateAt:      now.Add(-22 * time.Hour),
		},
		{
			SrcSitesID:    1,
			Status:        Enums.ProcessingCrawlLogStatus,
			TotalRequests: 1500,
			Requests:      1000,
			Success:       450,
			Failed:        550,
			RAMUsed:       25600,
			CPUUsed:       320,
			StartTime:     now.Add(-72 * time.Hour),
			EndTime:       &endTime,
			CreatedAt:     now.Add(-72 * time.Hour),
			UpdateAt:      now.Add(-2 * time.Hour),
		},
	}

	for _, item := range items {
		Facades.CrawlLogRepo().Create(&item)
	}
}

var _ Lib.DbSeederInterface = &CrawlLogDbSeeder{}
