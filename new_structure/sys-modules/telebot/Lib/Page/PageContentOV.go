package Page

import tele "gopkg.in/telebot.v4"

type PageContentOV struct {
	Message     string
	ReplyMarkup *tele.ReplyMarkup
	File        *tele.Document
	Photo       *tele.Photo
}
