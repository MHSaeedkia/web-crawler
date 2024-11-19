package Controllers

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	export "project-root/modules/export/Enums"
	PostEnums "project-root/modules/post/Enums"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type MainSelectedReportPage struct{}

func (p *MainSelectedReportPage) PageNumber() int {
	return Enums.MainSelectedReportPageNumber
}

func (p *MainSelectedReportPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnPost := newReplyMarkup.Data("Post", "btn_post")
	btnCreatExport := newReplyMarkup.Data("Create Export", "btn_crate_export")
	btnEdit := newReplyMarkup.Data("Edit", "btn_edit")
	btnDelete := newReplyMarkup.Data("Delete", "btn_delete")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnCreatExport, btnPost),
		newReplyMarkup.Row(btnEdit, btnDelete),
		newReplyMarkup.Row(btnBack),
	)
	report, _ := Facades.ReportRepo().FindReport(telSession.GetReportTempData().ReportIdSelected)
	message := fmt.Sprintf("report  “%s”  is selected, what operation do you want to perform?", report.Title)
	return &Page.PageContentOV{
		Message:     message,
		ReplyMarkup: newReplyMarkup,
	}
}

func (p *MainSelectedReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainSelectedReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_edit":
		return Page.GetPage(Enums.TitleUpdateReportPageNumber)
	case "btn_post":
		filter, _ := Facades.ReportFilterRepo().FindByReportId(telSession.GetReportTempData().ReportIdSelected)
		telSession.GetPostTempData().FilterId = filter.ID
		return Page.GetPage(PostEnums.MainPostSelectedReportPageNumber)
	case "btn_crate_export":
		filter, _ := Facades.ReportFilterRepo().FindByReportId(telSession.GetReportTempData().ReportIdSelected)
		telSession.GetPostTempData().FilterId = filter.ID
		return Page.GetPage(export.CreateExportReportPageNumber)
	case "btn_delete":
		err := Facades.ReportRepo().SoftDelete(telSession.GetReportTempData().ReportIdSelected)
		if err != nil {
			telSession.GetGeneralTempData().LastMessage = "delete error"
			return Page.GetPage(Enums.MainReportUserPageNumber)
		}
		telSession.GetReportTempData().ReportIdSelected = 0
		telSession.GetGeneralTempData().LastMessage = "delete successful"
		return Page.GetPage(Enums.MainReportUserPageNumber)
	}

	return StaticBtns.HandleIfClickBackBtn(btnKey, Enums.MainReportUserPageNumber)
}

var _ Page.PageInterface = &MainSelectedReportPage{}
