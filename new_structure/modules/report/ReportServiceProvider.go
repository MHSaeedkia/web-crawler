package user

import (
	"project-root/app"
	"project-root/modules/report/DB/Migrations"
	"project-root/modules/report/Http/Controllers"
	"project-root/modules/report/Http/Controllers/Create"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *ReportServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreateReportTable{})
	SysDatabase.RegisterMigration(&Migrations.CreateReportFilterTable{})

	// Pages
	Page.RegisterPages([]Page.PageInterface{
		&Controllers.MainReportUserPage{},

		// create
		&Create.TitleCreateReportPage{},
		&Create.IsNotificationCreateReportPage{},
	})
}

func (s *ReportServiceProvider) Boot() {

}

type ReportServiceProvider struct{}

var _ app.ServiceProviderInterface = &ReportServiceProvider{}
