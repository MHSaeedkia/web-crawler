package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreateUserTable struct{}

func (c CreateUserTable) Name() string {
	return "2024_11_7_1000_create_users_table"
}

func (c CreateUserTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.User{})
}

func (c CreateUserTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.User{})
}

var _ Lib.MigrationInterface = &CreateUserTable{}
