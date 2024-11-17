package StaticBtns

import (
	tele "gopkg.in/telebot.v4"
	"project-root/sys-modules/telebot/Lib/Page"
)

func GetBackStaticBtn() *tele.ReplyMarkup {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnBack),
	)
	return newReplyMarkup
}

func HandleIfClickBackBtn(btnKeyClicked string, destinationPageNumber int) Page.PageInterface {
	if btnKeyClicked == "btn_back" {
		return Page.GetPage(destinationPageNumber)
	}
	return nil
}
