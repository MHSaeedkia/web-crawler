package Range

import (
	"fmt"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"strconv"
	"strings"
)

type DealTypeUpdateFilterPage struct{}

func (p *DealTypeUpdateFilterPage) PageNumber() int {
	return Enums.DealTypeUpdateFilterPageNumber
}

func (p *DealTypeUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     GetDealTypesMessage(),
		ReplyMarkup: Filter.GetDefaultBtnUpdateFilter(),
	}
}

func GetDealTypesMessage() string {
	dealTypes := Enums.GetDealTypes()
	var sb strings.Builder
	sb.WriteString("Type your Deal type num:\n")
	for i, dealType := range dealTypes {
		sb.WriteString(fmt.Sprintf("%d => %s\n", i, dealType))
	}
	return sb.String()
}

func GetUserDealType(input string) (int, error) {
	dealTypes := Enums.GetDealTypes()
	num, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("invalid input, not a valid number")
	}

	if num < 0 || num >= len(dealTypes) {
		return -1, fmt.Errorf("invalid choice, please enter a number between 0 and %d", len(dealTypes)-1)
	}

	return num, nil
}

func (p *DealTypeUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {

	validInput, err := GetUserDealType(value)
	if err != nil {
		telSession.GetGeneralTempData().LastMessage = err.Error()
		return Page.GetPage(Enums.DealTypeUpdateFilterPageNumber)
	}

	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	filter.DealType = &validInput

	// --
	error := Facades.ReportFilterRepo().Update(filter)
	if error != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
		return Page.GetPage(Enums.DealTypeUpdateFilterPageNumber)
	}

	telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

func (p *DealTypeUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.DealType = nil
	})
}

var _ Page.PageInterface = &DealTypeUpdateFilterPage{}
