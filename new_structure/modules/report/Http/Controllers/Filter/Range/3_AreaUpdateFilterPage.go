package Range

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type AreaUpdateFilterPage struct{}

func (p *AreaUpdateFilterPage) PageNumber() int {
	return Enums.AreaUpdateFilterPageNumber
}

func (p *AreaUpdateFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return Filter.GetGeneratePageRangeUpdateFilter("Area")
}

func (p *AreaUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnInputRangeUpdateFilter(
		Enums.AreaUpdateFilterPageNumber,
		0,
		5000,
		value,
		telSession,
		func(num1, num2 *int, filter *Models2.ReportFilter) {
			filter.AreaMin = num1
			filter.AreaMax = num2
		})
}

func (p *AreaUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.AreaMin = nil
		filter.AreaMax = nil
	})
}

var _ Page.PageInterface = &AreaUpdateFilterPage{}
