package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/monitoring/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateCrawlLogTable struct{}

func (c CreateCrawlLogTable) Name() string {
	return "2024_11_18_1000_create_crawl_logs_table"
}

func (c CreateCrawlLogTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.CrawlLog{})
}

func (c CreateCrawlLogTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.CrawlLog{})
}

var _ Lib.MigrationInterface = &CreateCrawlLogTable{}
