package Lib

import "gorm.io/gorm"

type MigrationInterface interface {
	Name() string
	Up(db *gorm.DB)
	Down(db *gorm.DB)
}

var registeredMigrations = map[string]MigrationInterface{}

func RegisterMigration(migration MigrationInterface) {
	registeredMigrations[migration.Name()] = migration
}

func GetMigrations() map[string]MigrationInterface {
	return registeredMigrations
}
