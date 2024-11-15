package Page

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/user/DB/Models"
	"strconv"
)

type PageInterface interface {
	PageNumber() int
	GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup)
	OnInput(value string, telSession *Models.TelSession) PageInterface
	OnClickInlineBtn(btnKey string, telSession *Models.TelSession) PageInterface
}

var pageMap = make(map[int]PageInterface)

func RegisterPage(page PageInterface) {
	pageMap[page.PageNumber()] = page
}

func RegisterPages(pages []PageInterface) {
	for _, page := range pages {
		RegisterPage(page)
	}
}

func GetPage(pageNumber int) PageInterface {
	page, exists := pageMap[pageNumber]
	if !exists {
		panic("page not found: " + strconv.Itoa(pageNumber))
	}
	return page
}
