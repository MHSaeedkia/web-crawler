package Login

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/auth/Enums"
	"project-root/modules/auth/Http/Controllers/Register"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type UsernameLoginAuthPage struct{}

func (p *UsernameLoginAuthPage) PageNumber() int {
	return Enums.UsernameLoginAuthPageNumber
}

func (p *UsernameLoginAuthPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return "Login(1/2) - Please enter your username:", StaticBtns.GetBackStaticBtn()
}

func (p *UsernameLoginAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// need validation and check not exist on database
	telSession.GetAuthTempData().Username = value
	return Page.GetPage(Enums.PasswordLoginAuthPageNumber)
}

func (p *UsernameLoginAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainAuthPageNumber)
}

var _ Page.PageInterface = &Register.UsernameRegisterAuthPage{}
