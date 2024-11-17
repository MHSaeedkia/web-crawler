package Tests

import (
	"github.com/stretchr/testify/assert"
	"project-root/boot"
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	UserFacade "project-root/modules/user/Facades"
	"project-root/sys-modules/console/Lib"
	"testing"
)

func TestMain(m *testing.M) {
	boot.Bootstrap()
	Lib.CallManualCommand("database:migrate", []string{"up"})
	m.Run()
}

func TestTelSessionRepo(t *testing.T) {
	chatId := int64(123)

	// create
	_, err := UserFacade.TelSessionRepo().Create(&Models.TelSession{
		LoggedUserID: nil,
		ChatID:       chatId,
		CurrentPage:  Enums.MainAuthPageNumber,
		TempData:     map[string]interface{}{},
		//TempData:     map[string]interface{}{"key": "value"},
	})

	// find
	session, err := UserFacade.TelSessionRepo().FindByChatID(chatId)
	assert.NoError(t, err)
	assert.NotZero(t, session.ID)

	// update
	session.CurrentPage = 1
	updateErr := UserFacade.TelSessionRepo().UpdateByChatID(chatId, session)
	assert.NoError(t, updateErr)
	newSession, _ := UserFacade.TelSessionRepo().FindByChatID(chatId)
	assert.Equal(t, 1, newSession.CurrentPage)
}
