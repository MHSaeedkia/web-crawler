package Create

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type TitleUpdateReportPage struct{}

func (p *TitleUpdateReportPage) PageNumber() int {
	return Enums.TitleUpdateReportPageNumber
}

func (p *TitleUpdateReportPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnSkip := newReplyMarkup.Data("Skip", "btn_skip")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnSkip),
		newReplyMarkup.Row(btnBack),
	)
	report, _ := Facades.ReportRepo().FindReport(telSession.GetReportTempData().ReportIdSelected)
	telSession.GetReportTempData().Title = report.Title // for skip
	message := fmt.Sprintf("update Report(1/2) - Type a name for your report, this name must be unique:\n\n🔶 current value: %s", report.Title)
	return message, newReplyMarkup
}

func (p *TitleUpdateReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	// name validation
	if value == "" {
		telSession.GetGeneralTempData().LastMessage = "Title cannot be empty"
		return Page.GetPage(Enums.TitleUpdateReportPageNumber)
	}

	// duplicate name
	existItem, _ := Facades.ReportRepo().FindReportUserByTitle(*telSession.LoggedUserID, value)
	if existItem != nil {
		telSession.GetGeneralTempData().LastMessage = "You have already created a report with this title, you should not choose a duplicate name."
		return Page.GetPage(Enums.TitleUpdateReportPageNumber)
	}

	telSession.GetReportTempData().Title = value
	return Page.GetPage(Enums.IsNotificationUpdateReportPageNumber)
}

func (p *TitleUpdateReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	if btnKey == "btn_skip" {
		return Page.GetPage(Enums.IsNotificationUpdateReportPageNumber)
	}
	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainSelectedReportPageNumber)
}

var _ Page.PageInterface = &TitleUpdateReportPage{}
