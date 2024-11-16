package Range

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type ElevatorUpdateFilterPage struct{}

func (p *ElevatorUpdateFilterPage) PageNumber() int {
	return Enums.ElevatorUpdateFilterPageNumber
}

func (p *ElevatorUpdateFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return Filter.GetGeneratePageToggleUpdateFilter("Elevator")
}

func (p *ElevatorUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *ElevatorUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnClickInlineBtnToggleUpdateFilter(btnKey, telSession, func(isActive *int, filter *Models2.ReportFilter) {
		filter.Elevator = isActive
	})
}

var _ Page.PageInterface = &ElevatorUpdateFilterPage{}
