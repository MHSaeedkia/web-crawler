package user

import (
	"project-root/app"
	commands "project-root/modules/user/Commands"
	"project-root/modules/user/DB/Migrations"
	"project-root/modules/user/DB/Seeders"
	"project-root/modules/user/Http/Controllers"
	SysCommand "project-root/sys-modules/console/Lib"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *UserServiceProvider) Register() {
	// Migration
	SysDatabase.RegisterMigration(&Migrations.CreateUserTable{})
	SysDatabase.RegisterMigration(&Migrations.CreateTelSessionsTable{})

	// Commands
	SysCommand.RegisterCommand(&commands.MyCommand{})

	// Seeder
	SysDatabase.RegisterSeeders([]SysDatabase.DbSeederInterface{
		&Seeders.UserDbSeeder{},
	})

	// Pages
	Page.RegisterPages([]Page.PageInterface{
		// main
		&Controllers.MainUserPage{},
	})
}

func (s *UserServiceProvider) Boot() {

}

type UserServiceProvider struct{}

var _ app.ServiceProviderInterface = &UserServiceProvider{}
