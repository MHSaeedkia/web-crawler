package user

import (
	"project-root/app"
	"project-root/modules/export/DB/Migrations"
	"project-root/modules/export/DB/Seeders"
	"project-root/modules/export/Http/Controllers"
	"project-root/modules/export/Http/Controllers/Download"
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
	Page.RegisterPages([]Page.PageInterface{
		&Controllers.CreateExportReport{},
		&Controllers.MainExportPage{},
		&Download.SelectExportMethodPage{},
		&Download.SendFileExportInRobotMethodPage{},
	})
}

func (s *ExportServiceProvider) Boot() {

}

type ExportServiceProvider struct{}

var _ app.ServiceProviderInterface = &ExportServiceProvider{}
