package user

import (
	"project-root/app"
	"project-root/modules/source-site/DB/Migrations"
	"project-root/modules/source-site/DB/Seeders"
	SysDatabase "project-root/sys-modules/database/Lib"
)

func (s *SourceSiteServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreateSourceSiteTable{})

	// Seeders
	SysDatabase.RegisterSeeders([]SysDatabase.DbSeederInterface{
		&Seeders.SourceSiteDbSeeder{},
	})
}

func (s *SourceSiteServiceProvider) Boot() {

}

type SourceSiteServiceProvider struct{}

var _ app.ServiceProviderInterface = &SourceSiteServiceProvider{}
