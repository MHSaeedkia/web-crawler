package post

import (
	"project-root/app"
	"project-root/modules/post/DB/Migrations"
	"project-root/modules/post/DB/Seeders"
	"project-root/modules/post/Http/Controllers"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *PostServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreatePostTable{})

	// Seeders
	SysDatabase.RegisterSeeders([]SysDatabase.DbSeederInterface{
		&Seeders.PostDbSeeder{},
	})

	// Pages
	Page.RegisterPages([]Page.PageInterface{
		&Controllers.MainPostSelectedReportPage{},
		&Controllers.ShowSinglePostPage{},
	})
}

func (s *PostServiceProvider) Boot() {

}

type PostServiceProvider struct{}

var _ app.ServiceProviderInterface = &PostServiceProvider{}
