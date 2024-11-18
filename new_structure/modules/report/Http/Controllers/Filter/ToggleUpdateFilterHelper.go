package Filter

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

func getToggleBtnUpdateFilter() *tele.ReplyMarkup {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnYes := newReplyMarkup.Data("Yes", "btn_yes")
	btnNo := newReplyMarkup.Data("No", "btn_no")
	btnClear := newReplyMarkup.Data("Clear", "btn_clear")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnYes, btnNo),
		newReplyMarkup.Row(btnClear),
		newReplyMarkup.Row(btnBack),
	)
	return newReplyMarkup
}

func getMessageToggleUpdateFilter(name string) string {
	return "Update Report Filter - " + name + " has or not?"
}

func GetGeneratePageToggleUpdateFilter(name string) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     getMessageToggleUpdateFilter(name),
		ReplyMarkup: getToggleBtnUpdateFilter(),
	}
}

// TODO: refactor duplicate code
func OnClickInlineBtnToggleUpdateFilter(btnKey string, telSession *Models.TelSession, handleUpdate func(isActive *int, filter *Models2.ReportFilter)) Page.PageInterface {
	switch btnKey {
	case "btn_back":
		return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
	case "btn_clear":
		filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)

		// ---
		handleUpdate(nil, filter)

		// --
		error := Facades.ReportFilterRepo().Update(filter)
		if error != nil {
			telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
			return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
		}

		telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
		return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)

	case "btn_yes":
		filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
		isActive := 1
		handleUpdate(&isActive, filter)
		error := Facades.ReportFilterRepo().Update(filter)
		if error != nil {
			telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
			return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
		}

		telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
		return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)

	case "btn_no":
		filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
		isActive := 0
		handleUpdate(&isActive, filter)
		error := Facades.ReportFilterRepo().Update(filter)
		if error != nil {
			telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
			return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
		}

		telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
		return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
	}
	return nil
}
