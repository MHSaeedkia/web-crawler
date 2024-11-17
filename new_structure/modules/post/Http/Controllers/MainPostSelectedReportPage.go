package Controllers

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	PostModels "project-root/modules/post/DB/Models"
	PostEnums "project-root/modules/post/Enums"
	PostFacades "project-root/modules/post/Facades"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/Pars"
	Time "project-root/sys-modules/time/Lib"
	"strings"
)

type MainPostSelectedReportPage struct{}

func (p *MainPostSelectedReportPage) PageNumber() int {
	return PostEnums.MainPostSelectedReportPageNumber
}

func (p *MainPostSelectedReportPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}

	filter, _ := Facades.ReportFilterRepo().FindByReportId(telSession.GetReportTempData().ReportIdSelected)
	posts, countAllPage, err := PostFacades.PostRepo().GetPostsForFilter(filter, 1, telSession.GetPostTempData().LastPageNumber)

	var rows []tele.Row

	// dynamic btn
	for _, post := range *posts {
		btn := newReplyMarkup.Data(post.Title, fmt.Sprintf("btn_show_post_%d", post.ID))
		rows = append(rows, newReplyMarkup.Row(btn))
	}

	// static btn
	btnCreateNewReport := newReplyMarkup.Data("Create New Post", "btn_create_new_post")
	btnPreviousPage := newReplyMarkup.Data("previous page", "btn_previous_page")
	btnNextPage := newReplyMarkup.Data("next page", "btn_next_page")
	btnBack := newReplyMarkup.Data("Back", "btn_back")

	rows = append(rows, newReplyMarkup.Row(btnPreviousPage, btnNextPage))
	rows = append(rows, newReplyMarkup.Row(btnCreateNewReport))
	rows = append(rows, newReplyMarkup.Row(btnBack))

	newReplyMarkup.Inline(rows...)
	fmt.Println(countAllPage, err)
	return FormatPostList(posts, filter.Report.Title, countAllPage, telSession.GetPostTempData().LastPageNumber), newReplyMarkup
}

func FormatPostList(posts *[]PostModels.Post, titleReport string, allPages, currentPage int) string {
	if len(*posts) == 0 {
		return "No posts found."
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("The report \"%s\" has been selected, select a btn to view :\n\n", titleReport))

	for _, post := range *posts {
		link := fmt.Sprintf("(https://divar.ir/v/%s)", *post.ExternalSiteID)

		createdAt := Time.FormatTimeAgo(post.CreatedAt)
		updatedAt := "unknown"
		if post.UpdateAt != nil {
			updatedAt = Time.FormatTimeAgo(*post.UpdateAt)
		}

		price := "unknown"
		if post.Price != nil {
			price = fmt.Sprintf("%d", *post.Price)
		}

		result.WriteString(fmt.Sprintf(
			"%s\nprice:                         %s\npublic:                   %s\ncreated at:            %s\nupdated at:          %s\nlink:                         %s\n",
			post.Title,
			price,
			Pars.BoolParsForDisplay(post.IsPublic),
			createdAt,
			updatedAt,
			link,
		))
		result.WriteString("-----------------------------------\n")
	}

	formattedResult := strings.TrimSuffix(result.String(), "-----------------------------------\n")
	formattedResult += fmt.Sprintf("\n\nPage %d of %d", currentPage, allPages)

	return formattedResult
}

func (p *MainPostSelectedReportPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainPostSelectedReportPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_next_page":
		telSession.GetPostTempData().LastPageNumber += 1
		return Page.GetPage(PostEnums.MainPostSelectedReportPageNumber)
	case "btn_previous_page":
		if telSession.GetPostTempData().LastPageNumber > 1 {
			telSession.GetPostTempData().LastPageNumber -= 1
		} else {
			telSession.GetPostTempData().LastPageNumber = 1
		}
		return Page.GetPage(PostEnums.MainPostSelectedReportPageNumber)
	case "btn_create_new_report":
		//return Page.GetPage(ReportEnums.TitleCreateReportPageNumber)
	case "btn_back":
		return Page.GetPage(ReportEnums.MainSelectedReportPageNumber)
	default:
		// dynamic btn
		var itemIdID int
		if _, err := fmt.Sscanf(btnKey, "btn_show_post_%d", &itemIdID); err == nil {
			_, err := PostFacades.PostRepo().FindByID(itemIdID)
			if err != nil {
				telSession.GetGeneralTempData().LastMessage = "The post ID is not valid"
				return Page.GetPage(PostEnums.MainPostSelectedReportPageNumber)
			}
			telSession.GetPostTempData().PostId = itemIdID
			return Page.GetPage(PostEnums.ShowSinglePostPageNumber)
		}

	}
	return nil
}

var _ Page.PageInterface = &MainPostSelectedReportPage{}
