package Register

import (
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Facades"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type UsernameRegisterAuthPage struct{}

func (p *UsernameRegisterAuthPage) PageNumber() int {
	return Enums.UsernameRegisterAuthPageNumber
}

func (p *UsernameRegisterAuthPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     "Register(1/3) - Please enter your username:",
		ReplyMarkup: StaticBtns.GetBackStaticBtn(),
	}
}

func (p *UsernameRegisterAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// need validation and check not exist on database
	user, _ := Facades.UserRepo().FindByUsername(value)
	if user != nil {
		telSession.GetGeneralTempData().LastMessage = "This username has already been created"
		return Page.GetPage(Enums.UsernameRegisterAuthPageNumber)
	}

	// --
	telSession.GetAuthTempData().Username = value
	return Page.GetPage(Enums.PasswordRegisterAuthPageNumber)
}

func (p *UsernameRegisterAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainAuthPageNumber)
}

var _ Page.PageInterface = &UsernameRegisterAuthPage{}
