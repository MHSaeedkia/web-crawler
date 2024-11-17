package Range

import (
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type NeighborhoodUpdateFilterPage struct{}

func (p *NeighborhoodUpdateFilterPage) PageNumber() int {
	return Enums.NeighborhoodUpdateFilterPageNumber
}

func (p *NeighborhoodUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     "Type your Neighborhood name:",
		ReplyMarkup: Filter.GetDefaultBtnUpdateFilter(),
	}
}

func (p *NeighborhoodUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	filter.NeighborhoodName = &value

	// --
	error := Facades.ReportFilterRepo().Update(filter)
	if error != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
		return Page.GetPage(Enums.NeighborhoodUpdateFilterPageNumber)
	}

	telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

func (p *NeighborhoodUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.NeighborhoodName = nil
	})
}

var _ Page.PageInterface = &NeighborhoodUpdateFilterPage{}
