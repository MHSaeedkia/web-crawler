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
	pageContentOV := targetPage.GeneratePage(session)
	//message, replyMarkup := targetPage.GeneratePage(session)
	lastMessage := session.GetGeneralTempData().LastMessage
	if lastMessage != "" {
		pageContentOV.Message = lastMessage + "\n" + pageContentOV.Message
		session.GetGeneralTempData().LastMessage = ""
	}

	// 4_ update session
	UserFacade.TelSessionRepo().UpdateByChatID(chatId, session)

	// 5_ send message
	err := SendMessage(c, pageContentOV, isBtn)
	fmt.Println(err) // TODO: critical log
	return nil
}

func SendMessage(c tele.Context, pageContentOV *Page.PageContentOV, isBtn bool) error {
	if pageContentOV.Photo != nil {
		pageContentOV.Photo.Caption = pageContentOV.Message
		return sendOrEdit(c, pageContentOV.Photo, pageContentOV.ReplyMarkup, isBtn)
	}

	if pageContentOV.File != nil {
		pageContentOV.File.Caption = pageContentOV.Message
		return sendOrEdit(c, pageContentOV.File, pageContentOV.ReplyMarkup, isBtn)
	}

	return sendOrEdit(c, pageContentOV.Message, pageContentOV.ReplyMarkup, isBtn)
}

func sendOrEdit(c tele.Context, content interface{}, markup *tele.ReplyMarkup, isBtn bool) error {
	// check last message has media ? if click btn (Callback)
	lastMessageHasMedia := false
	if isBtn {
		msg := c.Callback().Message
		if msg.Photo != nil || msg.Document != nil {
			lastMessageHasMedia = true
		}
	}

	// remove last message if click btn and has media
	if isBtn && lastMessageHasMedia {
		c.Delete()
	}

	// update last message if last message is normal
	if isBtn && !lastMessageHasMedia {
		if markup != nil {
			return c.Edit(content, markup)
		}
		return c.Edit(content)
	}

	// send message if user not click btn
	// or click btn but last message has media
	if markup != nil {
		fmt.Println("HUUUUUUUasdaaaassssUUUUUUU", content)
		return c.Send(content, markup)
	}
	fmt.Println("HUUUUUUUUUUUUUU", content)
	return c.Send(content)
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
