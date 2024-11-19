package user

import (
	"project-root/app"
	Commands "project-root/modules/web-crawler/Commands"
	SysCommand "project-root/sys-modules/console/Lib"
)

func (s *WebCrawlerServiceProvider) Register() {
	// Commands
	SysCommand.RegisterCommand(&Commands.GetPostsWebCrawlerCommand{})
	SysCommand.RegisterCommand(&Commands.UpdatePricePostsWebCrawlerCommand{})
}

func (s *WebCrawlerServiceProvider) Boot() {

}

type WebCrawlerServiceProvider struct{}

var _ app.ServiceProviderInterface = &WebCrawlerServiceProvider{}
