package Controllers

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/export/DB/Models"
	"project-root/modules/export/Enums"
	Facades3 "project-root/modules/export/Facades"
	export "project-root/modules/export/Lib"
	"project-root/modules/post/Facades"
	ReportEnums "project-root/modules/report/Enums"
	Facades2 "project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"time"
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
		reportId := telSession.GetReportTempData().ReportIdSelected
		reportFilter, _ := Facades2.ReportFilterRepo().FindByReportId(reportId)
		posts, _, _ := Facades.PostRepo().GetPostsForFilter(reportFilter, 500, 1)

		// generate excel file
		fileName, err := export.FinalExport(*posts, Enums.ExcelFileType)

		if err != nil {
			telSession.GetGeneralTempData().LastMessage = "error while generate excel"
			return Page.GetPage(p.PageNumber())
		}

		// save export item
		createErr := Facades3.ExportRepo().Create(&Models2.Export{
			ReportID:  reportId,
			FileType:  Enums.ExcelFileType,
			FilePath:  fileName,
			CreatedAt: time.Now(),
		})
		if createErr != nil {
			telSession.GetGeneralTempData().LastMessage = "error while create db"
			return Page.GetPage(p.PageNumber())
		}

		telSession.GetGeneralTempData().LastMessage = "create xlsx successful"
		return Page.GetPage(ReportEnums.MainSelectedReportPageNumber)
	case "btn_csv":
		reportId := telSession.GetReportTempData().ReportIdSelected
		reportFilter, _ := Facades2.ReportFilterRepo().FindByReportId(reportId)
		posts, _, _ := Facades.PostRepo().GetPostsForFilter(reportFilter, 500, 1)

		// generate excel file
		fileName, err := export.FinalExport(*posts, Enums.CsvFileType)

		if err != nil {
			telSession.GetGeneralTempData().LastMessage = "error while generate csv"
			return Page.GetPage(p.PageNumber())
		}

		// save export item
		createErr := Facades3.ExportRepo().Create(&Models2.Export{
			ReportID:  reportId,
			FileType:  Enums.CsvFileType,
			FilePath:  fileName,
			CreatedAt: time.Now(),
		})
		if createErr != nil {
			telSession.GetGeneralTempData().LastMessage = "error while create db"
			return Page.GetPage(p.PageNumber())
		}

		telSession.GetGeneralTempData().LastMessage = "create csv successful"
		return Page.GetPage(ReportEnums.MainSelectedReportPageNumber)

	}
	return StaticBtns.HandleIfClickBackBtn(btnKey, ReportEnums.MainSelectedReportPageNumber)
}

var _ Page.PageInterface = &CreateExportReport{}
