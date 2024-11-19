package Download

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/export/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type SelectExportMethodPage struct{}

func (p *SelectExportMethodPage) PageNumber() int {
	return Enums.SelectExportMethodPageNumber
}

func (p *SelectExportMethodPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnEmail := newReplyMarkup.Data("Email", "btn_email")
	btnInRobot := newReplyMarkup.Data("In Robot", "btn_in_robot")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnEmail, btnInRobot),
		newReplyMarkup.Row(btnBack),
	)
	return &Page.PageContentOV{
		Message:     "Choose a method to send:",
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *SelectExportMethodPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *SelectExportMethodPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_email":
	case "btn_in_robot":
		/*		page := SendFileExportInRobotMethodPage{}
				pageContentOV := page.GeneratePage(telSession)
				Facades.Bot().Send(pageContentOV.File, pageContentOV.ReplyMarkup)*/
		return Page.GetPage(Enums.SendFileExportInRobotMethodPageNumber)
	}
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainExportPageNumber)
}

var _ Page.PageInterface = &SelectExportMethodPage{}
