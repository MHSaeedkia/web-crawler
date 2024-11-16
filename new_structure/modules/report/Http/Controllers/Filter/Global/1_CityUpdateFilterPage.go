package Range

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type CityUpdateFilterPage struct{}

func (p *CityUpdateFilterPage) PageNumber() int {
	return Enums.CityUpdateFilterPageNumber
}

func (p *CityUpdateFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return "Type your city name:", Filter.GetDefaultBtnUpdateFilter()
}

func (p *CityUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	filter.CityName = &value

	// --
	error := Facades.ReportFilterRepo().Update(filter)
	if error != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
		return Page.GetPage(Enums.CityUpdateFilterPageNumber)
	}

	telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

func (p *CityUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.CityName = nil
	})
}

var _ Page.PageInterface = &CityUpdateFilterPage{}
