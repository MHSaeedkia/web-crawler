package Create

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"strconv"
)

type IsNotificationUpdateReportPage struct{}

func (p *IsNotificationUpdateReportPage) PageNumber() int {
	return Enums.IsNotificationUpdateReportPageNumber
}

func (p *IsNotificationUpdateReportPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnYes := newReplyMarkup.Data("Yes", "btn_yes")
	btnNo := newReplyMarkup.Data("No", "btn_no")
	btnSkip := newReplyMarkup.Data("Skip", "btn_skip")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnNo, btnYes),
		newReplyMarkup.Row(btnSkip),
		newReplyMarkup.Row(btnBack),
	)

	report, _ := Facades.ReportRepo().FindReport(telSession.GetReportTempData().ReportIdSelected)
	telSession.GetReportTempData().IsNotification = report.IsNotification // for skip
	message := fmt.Sprintf("Update Report(2/2) - If it is checked again (watch-list), should the notification be sent to the robot?\n\n🔶 current value: %s", strconv.Itoa(report.IsNotification))
	return &Page.PageContentOV{
		Message:     message,
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *IsNotificationUpdateReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *IsNotificationUpdateReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	// duplicate name
	existItem, _ := Facades.ReportRepo().FindReportUserByTitle(*telSession.LoggedUserID, telSession.GetReportTempData().Title)
	if existItem != nil {
		telSession.GetGeneralTempData().LastMessage = "You have already created a report with this title, you should not choose a duplicate name."
		return Page.GetPage(Enums.TitleUpdateReportPageNumber)
	}

	// handle btn
	switch btnKey {
	case "btn_yes":
		return updateReport(1, telSession)
	case "btn_no":
		return updateReport(0, telSession)
	case "btn_skip":
		return updateReport(telSession.GetReportTempData().IsNotification, telSession)
	case "btn_back":
		return Page.GetPage(Enums.TitleUpdateReportPageNumber)
	}
	return nil
}

func updateReport(isNotification int, telSession *Models.TelSession) Page.PageInterface {
	report, errFindRep := Facades.ReportRepo().FindReport(telSession.GetReportTempData().ReportIdSelected)
	if errFindRep != nil {
		telSession.GetGeneralTempData().LastMessage = "report id not valid for update"
		return Page.GetPage(Enums.MainReportUserPageNumber)
	}
	report.Title = telSession.GetReportTempData().Title
	report.IsNotification = isNotification
	err := Facades.ReportRepo().Update(report)

	if err != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem while updating the report."
		return Page.GetPage(Enums.MainReportUserPageNumber)
	}

	filter, _ := Facades.ReportFilterRepo().FindByReportId(report.ID)

	telSession.GetReportTempData().ReportId = report.ID
	telSession.GetReportTempData().FilterId = filter.ID
	telSession.GetGeneralTempData().LastMessage = "The report was updated successfully."
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

var _ Page.PageInterface = &IsNotificationUpdateReportPage{}
