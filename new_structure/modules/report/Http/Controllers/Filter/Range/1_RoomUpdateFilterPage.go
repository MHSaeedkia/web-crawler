package Range

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type RoomUpdateFilterPage struct{}

func (p *RoomUpdateFilterPage) PageNumber() int {
	return Enums.RoomUpdateFilterPageNumber
}

func (p *RoomUpdateFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return Filter.GetGeneratePageRangeUpdateFilter("Room")
}

func (p *RoomUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnInputRangeUpdateFilter(
		Enums.RoomUpdateFilterPageNumber,
		1,
		10,
		value,
		telSession,
		func(num1, num2 *int, filter *Models2.ReportFilter) {
			filter.RoomCountMin = num1
			filter.RoomCountMax = num2
		})
}

func (p *RoomUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.RoomCountMin = nil
		filter.RoomCountMax = nil
	})
}

var _ Page.PageInterface = &RoomUpdateFilterPage{}
