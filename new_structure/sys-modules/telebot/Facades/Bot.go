package Facades

import (
	"gopkg.in/telebot.v4"
	"project-root/app"
)

func Bot() *telebot.Bot {
	return app.App.Resolve("bot").(*telebot.Bot)
}
