package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateTelSessionsTable struct{}

func (c CreateTelSessionsTable) Name() string {
	return "2024_11_10_1000_create_tel_sessions_table"
}

func (c CreateTelSessionsTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.TelSession{})
}

func (c CreateTelSessionsTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.TelSession{})
}

var _ Lib.MigrationInterface = &CreateTelSessionsTable{}
