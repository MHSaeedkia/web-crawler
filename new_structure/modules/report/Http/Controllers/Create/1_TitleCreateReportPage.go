package Create

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	UserEnums "project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type TitleCreateReportPage struct{}

func (p *TitleCreateReportPage) PageNumber() int {
	return Enums.TitleCreateReportPageNumber
}

func (p *TitleCreateReportPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return "Create Report(1/2) - Type a name for your report, this name must be unique:", StaticBtns.GetBackStaticBtn()
}

func (p *TitleCreateReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// name validation
	if value == "" {
		telSession.GetGeneralTempData().LastMessage = "Title cannot be empty"
		return Page.GetPage(Enums.TitleCreateReportPageNumber)
	}

	// duplicate name
	existItem, _ := Facades.ReportRepo().FindReportUserByTitle(*telSession.LoggedUserID, value)
	if existItem != nil {
		telSession.GetGeneralTempData().LastMessage = "You have already created a report with this title, you should not choose a duplicate name."
		return Page.GetPage(Enums.TitleCreateReportPageNumber)
	}
	telSession.GetReportTempData().Title = value
	return Page.GetPage(Enums.IsNotificationCreateReportPageNumber)
}

func (p *TitleCreateReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, UserEnums.MainUserPageNumber)
}

var _ Page.PageInterface = &TitleCreateReportPage{}
