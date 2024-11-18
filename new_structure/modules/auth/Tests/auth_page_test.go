package Tests

import (
	"github.com/stretchr/testify/assert"
	"project-root/boot"
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Bootstrap()
	Lib.CallManualCommand("database:migrate", []string{"up"})
	m.Run()
}

func TestAuthPage(t *testing.T) {
	authPage := Page.GetPage(Enums.MainAuthPageNumber)
	// register step
	session := &Models.TelSession{
		LoggedUserID: nil,
		ChatID:       123,
		CurrentPage:  Enums.MainAuthPageNumber,
		TempData:     map[string]interface{}{},
		//TempData:     map[string]interface{}{"key": "value"},
	}
	//session.TempData["last_message"] = "Please choose from the buttons below:"
	//lastMessage, exists := session.TempData["last_message"]
	//if !exists {
	//	lastMessage = ""
	//}
	//fmt.Println("value", value, "exists", exists, "bind")
	redirectPage := authPage.OnClickInlineBtn("btn_register", session)
	assert.Equal(t, Enums.UsernameRegisterAuthPageNumber, redirectPage.PageNumber())
}
