package post

import (
	"project-root/app"
	"project-root/modules/post/DB/Migrations"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *PostServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreatePostTable{})

	// Pages
	Page.RegisterPages([]Page.PageInterface{})
}

func (s *PostServiceProvider) Boot() {

}

type PostServiceProvider struct{}

var _ app.ServiceProviderInterface = &PostServiceProvider{}
