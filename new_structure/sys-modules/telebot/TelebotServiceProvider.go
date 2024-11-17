package env

import (
	tele "gopkg.in/telebot.v4"
	"project-root/app"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/env"
	commands "project-root/sys-modules/telebot/Commands"
	"project-root/sys-modules/telebot/Lib/Drivers/Default"
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

	app.App.Singleton("tele_bot_handle_input_driver", func() interface{} {
		return &Default.DefaultTeleBotHandleInputDriver{}
	})

}

type TelebotServiceProvider struct{}

var _ app.ServiceProviderInterface = &TelebotServiceProvider{}
