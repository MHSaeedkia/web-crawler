package boot

import (
	//"project-root/modules/tel-session"
	"fmt"
	"os"
	"project-root/app"
	"project-root/modules/auth"
	export "project-root/modules/export"
	monitoring "project-root/modules/monitoring"
	post "project-root/modules/post"
	report "project-root/modules/report"
	sourceSite "project-root/modules/source-site"
	"project-root/modules/user"
	webCrawler "project-root/modules/web-crawler"
	SysConsole "project-root/sys-modules/console"
	"project-root/sys-modules/console/Lib"
	SysDatabase "project-root/sys-modules/database"
	SysEnv "project-root/sys-modules/env"
	SysTelebot "project-root/sys-modules/telebot"
)

func Bootstrap() {
	modules := []app.ServiceProviderInterface{
		// ------------------ Modules -----------
		// -- sys
		&SysEnv.EnvServiceProvider{},
		&SysConsole.ConsoleServiceProvider{},
		&SysDatabase.DatabaseServiceProvider{},
		&SysTelebot.TelebotServiceProvider{},

		// -- app
		&user.UserServiceProvider{},
		&auth.AuthServiceProvider{},
		&report.ReportServiceProvider{},
		&sourceSite.SourceSiteServiceProvider{},
		&post.PostServiceProvider{},
		&webCrawler.WebCrawlerServiceProvider{},
		&monitoring.MonitoringServiceProvider{},
		&export.ExportServiceProvider{},
		// ---------------------------------------
	}

	// Register Modules
	for _, module := range modules {
		module.Register()
	}

	// Boot Modules
	for _, module := range modules {
		module.Boot()
	}
}

func HandleCommand() {
	if len(os.Args) < 2 {
		fmt.Println("Command not provided")
		return
	}

	commandName := os.Args[1]
	currentCommand, exist := Lib.GetCommand(commandName)
	if exist {
		currentCommand.Handle(os.Args[2:])
		return
	}
	println("Command not found with this signature!")
	return
}
