package commands

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"project-root/modules/auth/Enums"
	"project-root/modules/user/DB/Models"
	UserFacade "project-root/modules/user/Facades"
	"project-root/sys-modules/console/Lib"
	"project-root/sys-modules/telebot/Facades"
	"project-root/sys-modules/telebot/Lib/Page"
)

type StartTelebotCommand struct{}

func (c *StartTelebotCommand) Signature() string {
	return "telebot:start"
}

func (c *StartTelebotCommand) Description() string {
	return "Start and serve bot server"
}

func (c *StartTelebotCommand) Handle(args []string) {
	Facades.Bot().Handle("/start", func(c tele.Context) error {
		value := c.Text()
		return handleInput(false, value, c)
	})

	Facades.Bot().Handle(tele.OnText, func(c tele.Context) error {
		value := c.Text()
		return handleInput(false, value, c)
	})

	Facades.Bot().Handle(tele.OnCallback, func(c tele.Context) error {
		btnKey := c.Callback().Data[1:]
		return handleInput(true, btnKey, c)
	})

	startServer()
}

func handleInput(isBtn bool, value string, c tele.Context) error {
	// 1_ init var
	chatId := c.Sender().ID
	session := findSession(chatId)
	currentPage := Page.GetPage(session.CurrentPage)

	// 2_ call event
	targetPage := Page.GetPage(Enums.MainAuthPageNumber)
	if isBtn {
		targetPage = currentPage.OnClickInlineBtn(value, session)
	} else {
		targetPage = currentPage.OnInput(value, session)
	}

	// 404 page - back to currentPage
	if targetPage == nil {
		//c.Delete()
		//return nil
		targetPage = currentPage
		session.GetGeneralTempData().LastMessage = "Unknown behavior Please try again.\n"
	}

	// change page number
	session.CurrentPage = targetPage.PageNumber()

	// 3_ generate new message
	message, replyMarkup := targetPage.GeneratePage(session)
	lastMessage := session.GetGeneralTempData().LastMessage
	if lastMessage != "" {
		message = lastMessage + "\n" + message
		session.GetGeneralTempData().LastMessage = ""
	}

	// 4_ update session
	UserFacade.TelSessionRepo().UpdateByChatID(chatId, session)

	// 5_ send message
	if isBtn {
		err := c.Edit(message, replyMarkup)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		c.Send(message, replyMarkup)
	}
	return nil
}

// find or create session
func findSession(chatId int64) *Models.TelSession {
	fmt.Println(chatId)
	session, err := UserFacade.TelSessionRepo().FindByChatID(chatId)
	if err == nil {
		return session
	}

	newSession, _ := UserFacade.TelSessionRepo().Create(&Models.TelSession{
		LoggedUserID: nil,
		ChatID:       chatId,
		CurrentPage:  Enums.MainAuthPageNumber,
		TempData:     map[string]interface{}{},
		//TempData:     map[string]interface{}{"key": "value"},
	})
	return newSession
}

func startServer() {
	fmt.Println("server start...")
	Facades.Bot().Start()
}

var _ Lib.CommandInterface = &StartTelebotCommand{}
