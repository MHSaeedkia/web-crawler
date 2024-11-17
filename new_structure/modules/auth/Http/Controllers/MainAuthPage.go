package Controllers

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type MainAuthPage struct{}

func (p *MainAuthPage) PageNumber() int {
	return Enums.MainAuthPageNumber
}

func (p *MainAuthPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnOk := newReplyMarkup.Data("Login", "btn_login")
	btnNo := newReplyMarkup.Data("Register", "btn_register")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnOk, btnNo),
	)
	return &Page.PageContentOV{
		Message:     "Select an operation:",
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *MainAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_register":
		return Page.GetPage(Enums.UsernameRegisterAuthPageNumber)
	case "btn_login":
		return Page.GetPage(Enums.UsernameLoginAuthPageNumber)
	}
	return nil
}

var _ Page.PageInterface = &MainAuthPage{}
