package Range

import (
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type FloorUpdateFilterPage struct{}

func (p *FloorUpdateFilterPage) PageNumber() int {
	return Enums.FloorUpdateFilterPageNumber
}

func (p *FloorUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return Filter.GetGeneratePageRangeUpdateFilter("Floor")
}

func (p *FloorUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnInputRangeUpdateFilter(
		Enums.FloorUpdateFilterPageNumber,
		0,
		10,
		value,
		telSession,
		func(num1, num2 *int, filter *Models2.ReportFilter) {
			filter.FloorCountMin = num1
			filter.FloorCountMax = num2
		})
}

func (p *FloorUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.FloorCountMin = nil
		filter.FloorCountMax = nil
	})
}

var _ Page.PageInterface = &FloorUpdateFilterPage{}
