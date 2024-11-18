package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/report/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateReportTable struct{}

func (c CreateReportTable) Name() string {
	return "2024_11_9_1000_create_reports_table"
}

func (c CreateReportTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.Report{})
}

func (c CreateReportTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.Report{})
}

var _ Lib.MigrationInterface = &CreateReportTable{}
