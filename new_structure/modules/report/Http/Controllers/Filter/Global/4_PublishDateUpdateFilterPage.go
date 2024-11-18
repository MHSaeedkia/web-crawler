package Range

import (
	"fmt"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"strings"
	"time"
)

type PublishDateUpdateFilterPage struct{}

func (p *PublishDateUpdateFilterPage) PageNumber() int {
	return Enums.PublishDateUpdateFilterPageNumber
}

func (p *PublishDateUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return &Page.PageContentOV{
		Message:     "Enter the date range, so that first the start date, then the end date, separated by commas \n For example \n 2006-01-02 15:04:05,2008-02-01 00:00:00",
		ReplyMarkup: Filter.GetDefaultBtnUpdateFilter(),
	}
}

func (p *PublishDateUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {

	dateStart, dateEnd, err := ValidateAndParseDateRange(value)
	if err != nil {
		telSession.GetGeneralTempData().LastMessage = err.Error()
		return Page.GetPage(Enums.PublishDateUpdateFilterPageNumber)
	}

	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	filter.PostStartDate = &dateStart
	filter.PostEndDate = &dateEnd

	// --
	error := Facades.ReportFilterRepo().Update(filter)
	if error != nil {
		telSession.GetGeneralTempData().LastMessage = "There was a problem updating the filter"
		return Page.GetPage(Enums.PublishDateUpdateFilterPageNumber)
	}

	telSession.GetGeneralTempData().LastMessage = "The update was completed successfully"
	return Page.GetPage(Enums.MainUpdateReportFilterPageNumber)
}

func ValidateAndParseDateRange(input string) (time.Time, time.Time, error) {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("could not load timezone: %v", err)
	}

	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid input, must contain exactly two dates separated by a comma")
	}

	dateFormat := "2006-01-02 15:04:05"

	dateStart, err := time.ParseInLocation(dateFormat, strings.TrimSpace(parts[0]), loc)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid first date format, expected format is 'YYYY-MM-DD HH:MM:SS'")
	}

	dateEnd, err := time.ParseInLocation(dateFormat, strings.TrimSpace(parts[1]), loc)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid second date format, expected format is 'YYYY-MM-DD HH:MM:SS'")
	}

	if dateEnd.Before(dateStart) {
		return time.Time{}, time.Time{}, fmt.Errorf("the second date must be later than the first date")
	}

	return dateStart, dateEnd, nil
}

func (p *PublishDateUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.DefaultOnClickInlineBtnUpdateFilter(btnKey, telSession, func(filter *Models2.ReportFilter) {
		filter.PostStartDate = nil
		filter.PostEndDate = nil
	})
}

var _ Page.PageInterface = &PublishDateUpdateFilterPage{}
