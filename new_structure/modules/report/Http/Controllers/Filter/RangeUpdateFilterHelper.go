package Filter

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/report/Http/Requests"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

func GetDefaultBtnUpdateFilter() *tele.ReplyMarkup {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnClear := newReplyMarkup.Data("Clear", "btn_clear")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnClear),
		newReplyMarkup.Row(btnBack),
	)
	return newReplyMarkup
}

func getMessageRangeUpdateFilter(name string) string {
	return "Update Report Filter - Enter your " + name + " range, separated by \",\":\nFor example \"1,5\""
}

func GetGeneratePageRangeUpdateFilter(name string) (string, *tele.ReplyMarkup) {
	return getMessageRangeUpdateFilter(name), GetDefaultBtnUpdateFilter()
}

func OnInputRangeUpdateFilter(pageNum int, min int, max int, value string, telSession *Models.TelSession, handleUpdate func(num1, num2 *int, filter *Models2.ReportFilter)) Page.PageInterface {
	err, message, num1, num2 := Requests.ParsAndValidateRangeNumbers(value, min, max)
	if err {
		telSession.GetGeneralTempData().LastMessage = message
		return Page.GetPage(pageNum)
	}
	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	handleUpdate(&num1, &num2, filter)

	// --
	error := Facades.ReportFilterRepo().Update(filter)
	if error != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
		return Page.GetPage(pageNum)
	}

	telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

func DefaultOnClickInlineBtnUpdateFilter(btnKey string, telSession *Models.TelSession, handleClear func(filter *Models2.ReportFilter)) Page.PageInterface {
	switch btnKey {
	case "btn_back":
		return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
	case "btn_clear":
		filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)

		// ---
		handleClear(filter)

		// --
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
