package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/source-site/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateSourceSiteTable struct{}

func (c CreateSourceSiteTable) Name() string {
	return "2024_11_16_900_create_source_sites_table"
}

func (c CreateSourceSiteTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.SourceSite{})
}

func (c CreateSourceSiteTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.SourceSite{})
}

var _ Lib.MigrationInterface = &CreateSourceSiteTable{}
