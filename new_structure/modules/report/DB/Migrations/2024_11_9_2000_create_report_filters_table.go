package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateReportFilterTable struct{}

func (c CreateReportFilterTable) Name() string {
	return "2024_11_9_2000_create_report_filters_table"
}

func (c CreateReportFilterTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.ReportFilter{})
}

func (c CreateReportFilterTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.ReportFilter{})
}

var _ Lib.MigrationInterface = &CreateReportFilterTable{}
