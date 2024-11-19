package user

import (
	"project-root/app"
	"project-root/modules/export/DB/Migrations"
	"project-root/modules/export/DB/Seeders"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *ExportServiceProvider) Register() {
	// Migration
	SysDatabase.RegisterMigration(&Migrations.CreateExportTable{})

	// Seeder
	SysDatabase.RegisterSeeders([]SysDatabase.DbSeederInterface{
		&Seeders.ExportDbSeeder{},
	})

	// Pages
	Page.RegisterPages([]Page.PageInterface{})
}

func (s *ExportServiceProvider) Boot() {

}

type ExportServiceProvider struct{}

var _ app.ServiceProviderInterface = &ExportServiceProvider{}
