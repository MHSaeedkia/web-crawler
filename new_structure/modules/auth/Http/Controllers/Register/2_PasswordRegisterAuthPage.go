package Register

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type PasswordRegisterAuthPage struct{}

func (p *PasswordRegisterAuthPage) PageNumber() int {
	return Enums.PasswordRegisterAuthPageNumber
}

func (p *PasswordRegisterAuthPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return "Register(2/3) - Please enter your password:", StaticBtns.GetBackStaticBtn()
}

func (p *PasswordRegisterAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	telSession.GetAuthTempData().Password = value
	return Page.GetPage(Enums.EmailRegisterAuthPageNumber)
}

func (p *PasswordRegisterAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.UsernameRegisterAuthPageNumber)
}

var _ Page.PageInterface = &PasswordRegisterAuthPage{}
