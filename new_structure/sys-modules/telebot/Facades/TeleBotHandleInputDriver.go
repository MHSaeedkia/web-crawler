package Facades

import (
	"project-root/app"
	"project-root/sys-modules/telebot/Lib/Drivers"
)

func TeleBotHandleInputDriver() Drivers.TeleBotHandleInputDriverInterface {
	return app.App.Resolve("tele_bot_handle_input_driver").(Drivers.TeleBotHandleInputDriverInterface)
}
