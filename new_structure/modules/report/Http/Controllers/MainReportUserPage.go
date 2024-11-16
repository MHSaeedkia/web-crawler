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
	"strings"
	"time"
)

type MainReportUserPage struct{}

func (p *MainReportUserPage) PageNumber() int {
	return ReportEnums.MainReportUserPageNumber
}

func (p *MainReportUserPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}

	reports, countAllPage, err := Facades.ReportRepo().GetReportsByUserIdWithPagination(*telSession.LoggedUserID, 2, telSession.GetReportTempData().LastPageNumber)

	var rows []tele.Row

	// dynamic btn
	for _, report := range *reports {
		btn := newReplyMarkup.Data(report.Title, fmt.Sprintf("btn_show_report_%d", report.ID))
		rows = append(rows, newReplyMarkup.Row(btn))
	}

	// static btn
	btnCreateNewReport := newReplyMarkup.Data("Create New Report", "btn_create_new_report")
	btnPreviousPage := newReplyMarkup.Data("previous page", "btn_previous_page")
	btnNextPage := newReplyMarkup.Data("next page", "btn_next_page")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	rows = append(rows, newReplyMarkup.Row(btnPreviousPage, btnNextPage))
	rows = append(rows, newReplyMarkup.Row(btnCreateNewReport))
	rows = append(rows, newReplyMarkup.Row(btnBack))

	newReplyMarkup.Inline(rows...)
	fmt.Println(countAllPage, err)
	return FormatReportList(*reports, countAllPage, telSession.GetReportTempData().LastPageNumber), newReplyMarkup
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
	switch btnKey {
	case "btn_next_page":
		telSession.GetReportTempData().LastPageNumber += 1
		return Page.GetPage(ReportEnums.MainReportUserPageNumber)
	case "btn_previous_page":
		if telSession.GetReportTempData().LastPageNumber > 1 {
			telSession.GetReportTempData().LastPageNumber -= 1
		} else {
			telSession.GetReportTempData().LastPageNumber = 1
		}
		return Page.GetPage(ReportEnums.MainReportUserPageNumber)
	case "btn_create_new_report":
		return Page.GetPage(ReportEnums.TitleCreateReportPageNumber)
	case "btn_back":
		return Page.GetPage(Enums.MainUserPageNumber)
	default:
		// dynamic btn
		var reportID int
		if _, err := fmt.Sscanf(btnKey, "btn_show_report_%d", &reportID); err == nil {
			_, err := Facades.ReportRepo().FindReport(reportID)
			if err != nil {
				telSession.GetGeneralTempData().LastMessage = "The report ID is not valid"
				return Page.GetPage(ReportEnums.MainReportUserPageNumber)
			}
			telSession.GetReportTempData().ReportIdSelected = reportID
			return Page.GetPage(ReportEnums.MainSelectedReportPageNumber)
		}

	}
	return nil
}

var _ Page.PageInterface = &MainReportUserPage{}
