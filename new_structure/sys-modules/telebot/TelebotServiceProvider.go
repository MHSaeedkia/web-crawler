package env

import (
	tele "gopkg.in/telebot.v4"
	"project-root/app"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/env"
	commands "project-root/sys-modules/telebot/Commands"
	"time"
)

func (s *TelebotServiceProvider) Register() {
	Lib.RegisterCommand(&commands.StartTelebotCommand{})
	bindBot()
}

func (s *TelebotServiceProvider) Boot() {

}

func bindBot() {
	app.App.Singleton("bot", func() interface{} {
		pref := tele.Settings{
			Token:  env.Env("BOT_TOKEN"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}

		b, err := tele.NewBot(pref)
		if err != nil {
			panic(err)
		}
		return b
	})
}

type TelebotServiceProvider struct{}

var _ app.ServiceProviderInterface = &TelebotServiceProvider{}
