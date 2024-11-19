package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
	"project-root/modules/report/Facades"
	"project-root/sys-modules/database/Lib"
	"time"
)

type ReportDbSeeder struct{}

func (s ReportDbSeeder) Name() string {
	return "150_exports_table"
}

func (s ReportDbSeeder) Handle(db *gorm.DB) {
	userId := 2
	items := []Models.Report{
		{
			UserID:         &userId,
			Title:          "simple report",
			IsNotification: 1,
			CreatedAt:      time.Now(),
		},
	}

	for _, item := range items {
		Facades.ReportRepo().Create(&item)
	}

	filter := Models.ReportFilter{
		ReportID: 1,
	}
	Facades.ReportFilterRepo().Create(&filter)
}

var _ Lib.DbSeederInterface = &ReportDbSeeder{}
