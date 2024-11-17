package Login

import (
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	UserEnums "project-root/modules/user/Enums"
	"project-root/modules/user/Facades"
	"project-root/sys-modules/hash/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type PasswordLoginAuthPage struct{}

func (p *PasswordLoginAuthPage) PageNumber() int {
	return Enums.PasswordLoginAuthPageNumber
}

func (p *PasswordLoginAuthPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     "Login(2/2) - Please enter your password:",
		ReplyMarkup: StaticBtns.GetBackStaticBtn(),
	}
}

func (p *PasswordLoginAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	telSession.GetAuthTempData().Password = value

	user, _ := Facades.UserRepo().FindByUsername(telSession.GetAuthTempData().Username)
	if user == nil {
		telSession.GetGeneralTempData().LastMessage = "Such a username does not exist in the system"
		return Page.GetPage(Enums.UsernameLoginAuthPageNumber)
	}
	isValidPassword := Lib.CompareHashStr(user.Password, telSession.GetAuthTempData().Password)
	if !isValidPassword {
		telSession.GetGeneralTempData().LastMessage = "The password is not correct"
		return Page.GetPage(Enums.PasswordLoginAuthPageNumber)
	}

	// login successful
	telSession.GetGeneralTempData().LastMessage = "Login was successful"
	telSession.LoggedUser = user
	telSession.LoggedUserID = &user.ID

	// temp set default value
	telSession.GetAuthTempData().Password = ""
	telSession.GetAuthTempData().Username = ""
	return Page.GetPage(UserEnums.MainUserPageNumber)
}

func (p *PasswordLoginAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.UsernameLoginAuthPageNumber)
}

var _ Page.PageInterface = &PasswordLoginAuthPage{}
