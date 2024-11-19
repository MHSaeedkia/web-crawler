package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/export/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateExportTable struct{}

func (c CreateExportTable) Name() string {
	return "2024_11_19_1000_create_exports_table"
}

func (c CreateExportTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.Export{})
}

func (c CreateExportTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.Export{})
}

var _ Lib.MigrationInterface = &CreateExportTable{}
