package Register

import (
	tele "gopkg.in/telebot.v4"
	AuthEnums "project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/modules/user/Facades"
	"project-root/sys-modules/hash/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"regexp"
	"time"
)

type EmailRegisterAuthPage struct{}

func (p *EmailRegisterAuthPage) PageNumber() int {
	return AuthEnums.EmailRegisterAuthPageNumber
}

func (p *EmailRegisterAuthPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return "Register(3/3) - Please enter your email to receive exports:", StaticBtns.GetBackStaticBtn()
}

func (p *EmailRegisterAuthPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// validation email
	telSession.GetAuthTempData().Email = value
	if !isValidEmail(telSession.GetAuthTempData().Email) {
		telSession.GetGeneralTempData().LastMessage = "The email is not valid"
		return Page.GetPage(AuthEnums.EmailRegisterAuthPageNumber)
	}

	// validation username
	user, _ := Facades.UserRepo().FindByUsername(telSession.GetAuthTempData().Username)
	if user != nil {
		telSession.GetGeneralTempData().LastMessage = "This username has already been created"
		return Page.GetPage(AuthEnums.UsernameRegisterAuthPageNumber)
	}

	// hash and validation password
	passwordHashed, err := Lib.HashStr(telSession.GetAuthTempData().Password)
	if err != nil {
		telSession.GetGeneralTempData().LastMessage = "The password is not valid"
		return Page.GetPage(AuthEnums.PasswordRegisterAuthPageNumber)
	}

	// store
	err = Facades.UserRepo().Create(&Models.User{
		IsActive:      true,
		Username:      telSession.GetAuthTempData().Username,
		Password:      passwordHashed,
		RoleType:      Enums.UserRoleTypeEnum,
		Email:         value,
		CreatedChatID: &telSession.ChatID,
		LastError:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		//LastError: time.Now(),
	})

	if err == nil {
		telSession.GetGeneralTempData().LastMessage = "Registration was successful"
	} else {
		telSession.GetGeneralTempData().LastMessage = "There was a problem registering, contact support."
	}

	// temp set default value
	telSession.GetAuthTempData().Password = ""
	telSession.GetAuthTempData().Username = ""
	telSession.GetAuthTempData().Email = ""

	return Page.GetPage(AuthEnums.MainAuthPageNumber)
}

func (p *EmailRegisterAuthPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, AuthEnums.PasswordRegisterAuthPageNumber)
}

func isValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

func registerUser(string2 string) {
	Facades.UserRepo().Create(&Models.User{
		IsActive:      true,
		Username:      "admin",
		Password:      "123",
		RoleType:      Enums.AdminRoleTypeEnum,
		Email:         "info@foo.com",
		CreatedChatID: nil,
		LastError:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		//LastError: time.Now(),
	})
}

var _ Page.PageInterface = &PasswordRegisterAuthPage{}
