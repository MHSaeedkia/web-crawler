package Controllers

import (
	tele "gopkg.in/telebot.v4"
	"project-root/modules/export/Enums"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type CreateExportReport struct{}

func (p *CreateExportReport) PageNumber() int {
	return Enums.CreateExportReportPage
}

func (p *CreateExportReport) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnXlsx := newReplyMarkup.Data("xlsx", "btn_xlsx")
	btnCsv := newReplyMarkup.Data("csv", "btn_csv")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnXlsx, btnCsv),
		newReplyMarkup.Row(btnBack),
	)
	return &Page.PageContentOV{
		Message:     "Select the export file format:",
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *CreateExportReport) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *CreateExportReport) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_xlsx":
		telSession.GetGeneralTempData().LastMessage = "btn_xlsx"
		return Page.GetPage(p.PageNumber())
	case "btn_csv":
		telSession.GetGeneralTempData().LastMessage = "btn_csv"
		return Page.GetPage(p.PageNumber())
	}
	return StaticBtns.HandleIfClickBackBtn(btnKey, ReportEnums.MainSelectedReportPageNumber)
}

var _ Page.PageInterface = &CreateExportReport{}
