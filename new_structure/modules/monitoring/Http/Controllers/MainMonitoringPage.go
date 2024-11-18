package Controllers

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/monitoring/DB/Models"
	monitoringEnum "project-root/modules/monitoring/Enums"
	Facades2 "project-root/modules/monitoring/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"project-root/sys-modules/time/Lib"
)

type MainMonitoringPage struct{}

func (p *MainMonitoringPage) PageNumber() int {
	return monitoringEnum.MainUserMonitoringPageNumber
}

func (p *MainMonitoringPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {

	crawlLogs, countAllPage, _ := Facades2.CrawlLogRepo().GetCrawlLogsWithPagination(StaticBtns.GetDefaultPerPage(), telSession.GetMonitoringTempData().LastPageNumber)

	// dynamic btn
	paginationReplyMarkup := StaticBtns.PaginationReplyMarkupData{
		Items:        []StaticBtns.PaginationReplyMarkupItem{},
		StaticRowBtn: []tele.Row{},
	}

	// custom btn
	return &Page.PageContentOV{
		Message:     GenerateCrawlLogsToString(*crawlLogs, telSession.GetMonitoringTempData().LastPageNumber, countAllPage),
		ReplyMarkup: paginationReplyMarkup.GetReplyMarkup("usermonitoring"),
	}
}

func GenerateCrawlLogsToString(logs []Models2.CrawlLog, currentPage, pageSize int) string {
	var report string
	report += "crawl_logs - activities:\n"

	for _, log := range logs {
		startedAgo := Lib.FormatTimeAgo(log.StartTime)
		lastUpdatedAgo := Lib.FormatTimeAgo(log.UpdateAt)
		endTime := "-"
		if log.EndTime != nil {
			endTime = Lib.FormatTimeAgo(*log.EndTime)
		}

		statusLabel := monitoringEnum.GetCrawlLogStatusLabel(log.Status)

		report += fmt.Sprintf(
			"Started %s\n"+
				"end time:              %s\n"+
				"status:                  %s\n"+
				"req to be sent:   %d\n"+
				"send req:             %d\n"+
				"fail req:                 %d\n"+
				"successful req:  %d\n"+
				"ram_used:            %dmb\n"+
				"cpu used:             %d\n"+
				"last update:         %s\n"+
				"-----------------------------------\n",
			startedAgo,
			endTime,
			statusLabel,
			log.TotalRequests,
			log.Requests,
			log.Failed,
			log.Success,
			log.RAMUsed,
			log.CPUUsed,
			lastUpdatedAgo,
		)
	}

	report += fmt.Sprintf("Page %d of %d\n", currentPage, pageSize)
	return report
}

func (p *MainMonitoringPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainMonitoringPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	// pagination btn
	paginationHandleBtn := StaticBtns.PaginationOnClickBtnData{
		PrefixBtnKey:      "usermonitoring",
		CurrentPageNumber: p.PageNumber(),
		BackPageNumber:    Enums.MainUserPageNumber,
		GetPageNumberSaved: func() int {
			return telSession.GetMonitoringTempData().LastPageNumber
		},
		SavePageNumber: func(pageNum int) {
			telSession.GetMonitoringTempData().LastPageNumber = pageNum
		},
		OnSelectItemId: func(itemId int) Page.PageInterface {
			return nil
		},
	}
	return paginationHandleBtn.HandleInputPagination(btnKey)
}

var _ Page.PageInterface = &MainMonitoringPage{}
