package Migrations

import (
	"gorm.io/gorm"
	"project-root/modules/post/DB/Models"
	"project-root/sys-modules/database/Lib"
)

type CreatePostTable struct{}

func (c CreatePostTable) Name() string {
	return "2024_11_16_1000_create_posts_table"
}

func (c CreatePostTable) Up(db *gorm.DB) {
	db.Migrator().CreateTable(&Models.Post{})
}

func (c CreatePostTable) Down(db *gorm.DB) {
	db.Migrator().DropTable(&Models.Post{})
}

var _ Lib.MigrationInterface = &CreatePostTable{}
