package Facades

import (
	"gorm.io/gorm"
	"project-root/app"
)

func Db() *gorm.DB {
	return app.App.Resolve("db_connection").(*gorm.DB)
}
