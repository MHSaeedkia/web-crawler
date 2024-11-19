package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/export/DB/Models"
	Enums2 "project-root/modules/export/Enums"
	"project-root/modules/export/Facades"
	"project-root/sys-modules/database/Lib"
	"time"
)

type ExportDbSeeder struct{}

func (s ExportDbSeeder) Name() string {
	return "500_exports_table"
}

func (s ExportDbSeeder) Handle(db *gorm.DB) {
	Facades.ExportRepo().Truncate()
	now := time.Now()
	items := []Models.Export{
		{
			ReportID:  1,
			FileType:  Enums2.ExcelFileType,
			FilePath:  "",
			CreatedAt: now.Add(-92 * time.Hour),
		},
		{
			ReportID:  1,
			FileType:  Enums2.CsvFileType,
			FilePath:  "",
			CreatedAt: now.Add(-5 * time.Hour),
		},
	}

	for _, item := range items {
		Facades.ExportRepo().Create(&item)
	}
}

var _ Lib.DbSeederInterface = &ExportDbSeeder{}
