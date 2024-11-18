package Controllers

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	ReportModels "project-root/modules/report/DB/Models"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/modules/user/Enums"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"strings"
	"time"
)

type MainReportUserPage struct{}

func (p *MainReportUserPage) PageNumber() int {
	return ReportEnums.MainReportUserPageNumber
}

func (p *MainReportUserPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {

	reports, countAllPage, _ := Facades.ReportRepo().GetReportsByUserIdWithPagination(
		*telSession.LoggedUserID,
		StaticBtns.GetDefaultPerPage(),
		telSession.GetReportTempData().LastPageNumber,
	)

	// dynamic btn
	paginationReplyMarkup := StaticBtns.PaginationReplyMarkupData{
		Items:        []StaticBtns.PaginationReplyMarkupItem{},
		StaticRowBtn: []tele.Row{},
	}

	for _, item := range *reports {
		paginationReplyMarkup.Items = append(paginationReplyMarkup.Items, StaticBtns.PaginationReplyMarkupItem{
			ID:    item.ID,
			Title: item.Title,
		})
	}

	// custom btn
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnCreateNewReport := newReplyMarkup.Data("Create New Report", "btn_create_new_report")
	paginationReplyMarkup.StaticRowBtn = append(paginationReplyMarkup.StaticRowBtn, newReplyMarkup.Row(btnCreateNewReport))

	return &Page.PageContentOV{
		Message:     FormatReportList(*reports, countAllPage, telSession.GetReportTempData().LastPageNumber),
		ReplyMarkup: paginationReplyMarkup.GetReplyMarkup("report"),
	}
}

func FormatReportList(reports []ReportModels.Report, allPages int, currentPage int) string {
	if len(reports) == 0 {
		return "No reports found."
	}

	var result strings.Builder
	result.WriteString("Select a report to edit, delete or display the obtained results:\n\n")

	for _, report := range reports {
		notification := "❌"
		if report.IsNotification == 1 {
			notification = "✅"
		}
		var createAt string
		duration := time.Since(report.CreatedAt)
		if duration < time.Minute {
			createAt = "moments ago"
		} else if duration < time.Hour {
			createAt = fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
		} else if duration < 24*time.Hour {
			createAt = fmt.Sprintf("%d hours ago", int(duration.Hours()))
		} else {
			createAt = fmt.Sprintf("%d days ago", int(duration.Hours()/24))
		}

		result.WriteString(fmt.Sprintf(
			"%s\nnotification:         %s\ncreate at:              %s\n",
			report.Title,
			notification,
			createAt,
		))
		result.WriteString("-----------------------------------\n")
	}
	formattedResult := strings.TrimSuffix(result.String(), "-----------------------------------\n")

	formattedResult += fmt.Sprintf("\n\nPage %d of %d", currentPage, allPages)

	return formattedResult
}

func (p *MainReportUserPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainReportUserPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	// custom btn
	if btnKey == "btn_create_new_report" {
		return Page.GetPage(ReportEnums.TitleCreateReportPageNumber)
	}

	// pagination btn
	paginationHandleBtn := StaticBtns.PaginationOnClickBtnData{
		PrefixBtnKey:      "report",
		CurrentPageNumber: p.PageNumber(),
		BackPageNumber:    Enums.MainUserPageNumber,
		GetPageNumberSaved: func() int {
			return telSession.GetReportTempData().LastPageNumber
		},
		SavePageNumber: func(pageNum int) {
			telSession.GetReportTempData().LastPageNumber = pageNum
		},
		OnSelectItemId: func(itemId int) Page.PageInterface {
			_, err := Facades.ReportRepo().FindReport(itemId)
			if err != nil {
				telSession.GetGeneralTempData().LastMessage = "The report ID is not valid"
				return Page.GetPage(ReportEnums.MainReportUserPageNumber)
			}
			telSession.GetReportTempData().ReportIdSelected = itemId
			return Page.GetPage(ReportEnums.MainSelectedReportPageNumber)
		},
	}
	return paginationHandleBtn.HandleInputPagination(btnKey)
}

var _ Page.PageInterface = &MainReportUserPage{}
