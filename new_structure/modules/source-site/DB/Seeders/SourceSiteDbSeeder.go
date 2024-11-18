package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/source-site/DB/Models"
	"project-root/modules/source-site/Facades"
	"project-root/sys-modules/database/Lib"
)

type SourceSiteDbSeeder struct{}

func (s SourceSiteDbSeeder) Name() string {
	return "200_source_sites_table"
}

func (s SourceSiteDbSeeder) Handle(db *gorm.DB) {
	Facades.SourceSiteRepo().Truncate()
	sourceSite1 := &Models.SourceSite{
		ID:            1,
		Title:         "Divar",
		Config:        map[string]interface{}{"key": "value", "active": true},
		CrawlInterval: 10,
		MaxPosts:      100,
	}
	sourceSite2 := &Models.SourceSite{
		ID:            2,
		Title:         "Sheypoor",
		Config:        map[string]interface{}{"key": "value", "active": false},
		CrawlInterval: 10,
		MaxPosts:      100,
	}
	Facades.SourceSiteRepo().Create(sourceSite1)
	Facades.SourceSiteRepo().Create(sourceSite2)
}

var _ Lib.DbSeederInterface = &SourceSiteDbSeeder{}
