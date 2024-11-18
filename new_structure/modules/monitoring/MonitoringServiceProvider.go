package user

import (
	"project-root/app"
	"project-root/modules/monitoring/DB/Migrations"
	"project-root/modules/monitoring/DB/Seeders"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *MonitoringServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreateCrawlLogTable{})

	// Seeders
	SysDatabase.RegisterSeeders([]SysDatabase.DbSeederInterface{
		&Seeders.CrawlLogDbSeeder{},
	})

	// Pages
	Page.RegisterPages([]Page.PageInterface{})
}

func (s *MonitoringServiceProvider) Boot() {

}

type MonitoringServiceProvider struct{}

var _ app.ServiceProviderInterface = &MonitoringServiceProvider{}
