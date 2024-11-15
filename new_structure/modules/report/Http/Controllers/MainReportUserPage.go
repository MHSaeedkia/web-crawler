package Controllers

import (
	tele "gopkg.in/telebot.v4"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
)

type MainReportUserPage struct{}

func (p *MainReportUserPage) PageNumber() int {
	return ReportEnums.MainReportUserPageNumber
}

func (p *MainReportUserPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnCreateNewReport := newReplyMarkup.Data("Create New Report", "btn_create_new_report")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnCreateNewReport),
		newReplyMarkup.Row(btnBack),
	)
	return "report list ...", newReplyMarkup
}

func (p *MainReportUserPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainReportUserPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_create_new_report":
		return Page.GetPage(ReportEnums.TitleCreateReportPageNumber)
	case "btn_back":
		return Page.GetPage(Enums.MainUserPageNumber)
	}
	return nil
}

var _ Page.PageInterface = &MainReportUserPage{}
