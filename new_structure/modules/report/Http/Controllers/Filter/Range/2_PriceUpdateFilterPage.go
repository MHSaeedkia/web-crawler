package Range

import (
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type PriceUpdateFilterPage struct{}

func (p *PriceUpdateFilterPage) PageNumber() int {
	return Enums.PriceUpdateFilterPageNumber
}

func (p *PriceUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return Filter.GetGeneratePageRangeUpdateFilter("Price")
}

func (p *PriceUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnInputRangeUpdateFilter(
		Enums.PriceUpdateFilterPageNumber,
		0,
		100000000,
		value,
		telSession,
		func(num1, num2 *int, filter *Models2.ReportFilter) {
			filter.PriceMin = num1
			filter.PriceMax = num2
		})
}

func (p *PriceUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.PriceMin = nil
		filter.PriceMax = nil
	})
}

var _ Page.PageInterface = &PriceUpdateFilterPage{}
