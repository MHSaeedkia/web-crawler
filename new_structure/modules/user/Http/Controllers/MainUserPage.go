package Controllers

import (
	tele "gopkg.in/telebot.v4"
	AuthEnums "project-root/modules/auth/Enums"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
)

type MainUserPage struct{}

func (p *MainUserPage) PageNumber() int {
	return Enums.MainUserPageNumber
}

func (p *MainUserPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnBookmark := newReplyMarkup.Data("Bookmark", "btn_bookmark")
	btnMonitoring := newReplyMarkup.Data("Monitoring", "btn_monitoring")
	btnExport := newReplyMarkup.Data("Export", "btn_export")
	btnReport := newReplyMarkup.Data("Report", "btn_report")
	btnLogout := newReplyMarkup.Data("Logout", "btn_logout")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnExport, btnReport),
		newReplyMarkup.Row(btnBookmark, btnMonitoring),
		newReplyMarkup.Row(btnLogout),
	)
	return "Welcome dear " + telSession.LoggedUser.Username + " user", newReplyMarkup
}

func (p *MainUserPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainUserPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_bookmark":
	case "btn_monitoring":
	case "btn_export":
	case "btn_report":
		return Page.GetPage(ReportEnums.MainReportUserPageNumber)
	case "btn_logout":
		telSession.GetGeneralTempData().LastMessage = "You have successfully logged out"
		telSession.LoggedUserID = nil
		telSession.LoggedUser = nil
		return Page.GetPage(AuthEnums.MainAuthPageNumber)
	}
	return nil
}

var _ Page.PageInterface = &MainUserPage{}
