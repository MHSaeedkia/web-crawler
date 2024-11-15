package database

import (
	"gorm.io/gorm"
	"project-root/app"
	SysCommand "project-root/sys-modules/console/Lib"
	commands "project-root/sys-modules/database/Commands"
	"project-root/sys-modules/database/Lib/Connections"
	"project-root/sys-modules/env"
)

func (s *DatabaseServiceProvider) Register() {
	// connections - strategy pattern
	app.App.Singleton("mysql", func() interface{} {
		return Connections.NewMysqlDbConnection()
	})

	app.App.Singleton("sqlite", func() interface{} {
		return Connections.NewSqliteDbConnection()
	})

	app.App.Singleton("in_memory_sqlite", func() interface{} {
		return Connections.NewInMemorySqliteDbConnection()
	})

	app.App.Singleton("db_connection", func() interface{} {
		connectionName := env.Env("DB_CONNECTION")
		return app.App.ResolveWithoutLock(connectionName).(*gorm.DB)
	})

	// commands
	SysCommand.RegisterCommand(&commands.MigrationCommand{})
	SysCommand.RegisterCommand(&commands.SeedCommand{})
	SysCommand.RegisterCommand(&commands.FreshCommand{})
}

func (s *DatabaseServiceProvider) Boot() {

}

type DatabaseServiceProvider struct{}

var _ app.ServiceProviderInterface = &DatabaseServiceProvider{}
