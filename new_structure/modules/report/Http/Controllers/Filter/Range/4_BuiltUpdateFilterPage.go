package Range

import (
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type BuiltUpdateFilterPage struct{}

func (p *BuiltUpdateFilterPage) PageNumber() int {
	return Enums.BuiltUpdateFilterPageNumber
}

func (p *BuiltUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return Filter.GetGeneratePageRangeUpdateFilter("Built Year")
}

func (p *BuiltUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnInputRangeUpdateFilter(
		Enums.BuiltUpdateFilterPageNumber,
		1200,
		1500,
		value,
		telSession,
		func(num1, num2 *int, filter *Models2.ReportFilter) {
			filter.BuiltStart = num1
			filter.BuiltEnd = num2
		})
}

func (p *BuiltUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.BuiltStart = nil
		filter.BuiltEnd = nil
	})
}

var _ Page.PageInterface = &PriceUpdateFilterPage{}
