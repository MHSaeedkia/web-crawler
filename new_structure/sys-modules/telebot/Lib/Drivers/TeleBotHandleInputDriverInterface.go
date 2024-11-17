package Drivers

import tele "gopkg.in/telebot.v4"

type TeleBotHandleInputDriverInterface interface {
	HandleInput(isBtn bool, value string, c tele.Context) error
}
