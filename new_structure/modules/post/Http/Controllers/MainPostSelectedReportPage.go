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
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	Time "project-root/sys-modules/time/Lib"
	"strings"
)

type MainPostSelectedReportPage struct{}

func (p *MainPostSelectedReportPage) PageNumber() int {
	return PostEnums.MainPostSelectedReportPageNumber
}

func (p *MainPostSelectedReportPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	filter, _ := Facades.ReportFilterRepo().FindByReportId(telSession.GetReportTempData().ReportIdSelected)
	posts, countAllPage, _ := PostFacades.PostRepo().GetPostsForFilter(
		filter,
		StaticBtns.GetDefaultPerPage(),
		telSession.GetPostTempData().LastPageNumber,
	)

	// dynamic btn
	paginationReplyMarkup := StaticBtns.PaginationReplyMarkupData{
		Items:        []StaticBtns.PaginationReplyMarkupItem{},
		StaticRowBtn: []tele.Row{},
	}

	for _, post := range *posts {
		paginationReplyMarkup.Items = append(paginationReplyMarkup.Items, StaticBtns.PaginationReplyMarkupItem{
			ID:    post.ID,
			Title: post.Title,
		})
	}

	return &Page.PageContentOV{
		Message:     FormatPostList(posts, filter.Report.Title, countAllPage, telSession.GetPostTempData().LastPageNumber),
		ReplyMarkup: paginationReplyMarkup.GetReplyMarkup("post"),
	}
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
	paginationHandleBtn := StaticBtns.PaginationOnClickBtnData{
		PrefixBtnKey:      "post",
		CurrentPageNumber: p.PageNumber(),
		BackPageNumber:    ReportEnums.MainSelectedReportPageNumber,
		GetPageNumberSaved: func() int {
			return telSession.GetPostTempData().LastPageNumber
		},
		SavePageNumber: func(pageNum int) {
			telSession.GetPostTempData().LastPageNumber = pageNum
		},
		OnSelectItemId: func(itemId int) Page.PageInterface {
			_, err := PostFacades.PostRepo().FindByID(itemId)
			if err != nil {
				telSession.GetGeneralTempData().LastMessage = "The post ID is not valid"
				return Page.GetPage(PostEnums.MainPostSelectedReportPageNumber)
			}
			telSession.GetPostTempData().PostId = itemId
			return Page.GetPage(PostEnums.ShowSinglePostPageNumber)
		},
	}
	return paginationHandleBtn.HandleInputPagination(btnKey)
}

var _ Page.PageInterface = &MainPostSelectedReportPage{}
