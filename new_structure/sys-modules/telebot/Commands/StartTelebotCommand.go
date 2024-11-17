package commands

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/telebot/Facades"
)

type StartTelebotCommand struct{}

func (c *StartTelebotCommand) Signature() string {
	return "telebot:start"
}

func (c *StartTelebotCommand) Description() string {
	return "Start and serve bot server"
}

func (c *StartTelebotCommand) Handle(args []string) {
	Facades.Bot().Handle("/start", func(c tele.Context) error {
		value := c.Text()
		return Facades.TeleBotHandleInputDriver().HandleInput(false, value, c)
	})

	Facades.Bot().Handle(tele.OnText, func(c tele.Context) error {
		value := c.Text()
		return Facades.TeleBotHandleInputDriver().HandleInput(false, value, c)
	})

	Facades.Bot().Handle(tele.OnCallback, func(c tele.Context) error {
		btnKey := c.Callback().Data[1:]
		return Facades.TeleBotHandleInputDriver().HandleInput(true, btnKey, c)
	})

	startServer()
}

func startServer() {
	fmt.Println("server start...")
	Facades.Bot().Start()
}

var _ Lib.CommandInterface = &StartTelebotCommand{}
