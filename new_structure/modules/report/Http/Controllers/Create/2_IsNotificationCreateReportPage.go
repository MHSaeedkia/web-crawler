package Create

import (
	tele "gopkg.in/telebot.v4"
	ReportModels "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"time"
)

type IsNotificationCreateReportPage struct{}

func (p *IsNotificationCreateReportPage) PageNumber() int {
	return Enums.IsNotificationCreateReportPageNumber
}

func (p *IsNotificationCreateReportPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnYes := newReplyMarkup.Data("Yes", "btn_yes")
	btnNo := newReplyMarkup.Data("No", "btn_no")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnNo, btnYes),
		newReplyMarkup.Row(btnBack),
	)
	return &Page.PageContentOV{
		Message:     "Create Report(2/2) - If it is checked again (watch-list), should the notification be sent to the robot?",
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *IsNotificationCreateReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *IsNotificationCreateReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	// duplicate name
	existItem, _ := Facades.ReportRepo().FindReportUserByTitle(*telSession.LoggedUserID, telSession.GetReportTempData().Title)
	if existItem != nil {
		telSession.GetGeneralTempData().LastMessage = "You have already created a report with this title, you should not choose a duplicate name."
		return Page.GetPage(Enums.TitleCreateReportPageNumber)
	}

	// handle btn
	switch btnKey {
	case "btn_yes":
		return createReport(1, telSession)

	case "btn_no":
		return createReport(0, telSession)
	case "btn_back":
		return Page.GetPage(Enums.TitleCreateReportPageNumber)
	}
	return nil
}

func createReport(isNotification int, telSession *Models.TelSession) Page.PageInterface {
	report, err := Facades.ReportRepo().Create(&ReportModels.Report{
		UserID:         telSession.LoggedUserID,
		Title:          telSession.GetReportTempData().Title,
		IsNotification: isNotification,
		CreatedAt:      time.Now(),
	})

	if err != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem while creating the report."
		return Page.GetPage(Enums.MainReportUserPageNumber)
	}

	filter, errF := Facades.ReportFilterRepo().Create(&ReportModels.ReportFilter{
		ReportID: report.ID,
		Report:   report,
	})

	if errF != nil {
		Facades.ReportRepo().Delete(report.ID)
		telSession.GetGeneralTempData().LastMessage = "There was a problem while creating the report/filter."
		return Page.GetPage(Enums.MainReportUserPageNumber)
	}

	telSession.GetReportTempData().ReportId = report.ID
	telSession.GetReportTempData().FilterId = filter.ID
	telSession.GetGeneralTempData().LastMessage = "The report was created successfully."
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

var _ Page.PageInterface = &IsNotificationCreateReportPage{}
