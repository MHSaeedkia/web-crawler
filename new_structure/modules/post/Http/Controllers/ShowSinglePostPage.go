package Controllers

import (
	tele "gopkg.in/telebot.v4"
	PostEnums "project-root/modules/post/Enums"
	Facades2 "project-root/modules/post/Facades"
	"project-root/modules/post/Lib"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
)

type ShowSinglePostPage struct{}

func (p *ShowSinglePostPage) PageNumber() int {
	return PostEnums.ShowSinglePostPageNumber
}

func (p *ShowSinglePostPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnPriceHistory := newReplyMarkup.Data("Price history", "btn_price_history")
	btnBookmark := newReplyMarkup.Data("Bookmark", "btn_bookmark")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnBookmark, btnPriceHistory),
		newReplyMarkup.Row(btnBack),
	)
	post, _ := Facades2.PostRepo().FindByID(telSession.GetPostTempData().PostId)
	page := &Page.PageContentOV{
		Message:     Lib.FormatPostText(post),
		ReplyMarkup: newReplyMarkup,
	}

	if post != nil && post.MainIMG != nil && *post.MainIMG != "" {
		page.Photo = &tele.Photo{
			File: tele.FromURL(*post.MainIMG),
		}
	}
	return page

}

func (p *ShowSinglePostPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *ShowSinglePostPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {

	switch btnKey {
	case "btn_price_history":
		post, _ := Facades2.PostRepo().FindByID(telSession.GetPostTempData().PostId)
		if post.PriceHistory == nil || len(post.PriceHistory) <= 0 {
			telSession.GetGeneralTempData().LastMessage = "We have no price history"
			return Page.GetPage(p.PageNumber())
		}
		return Page.GetPage(PostEnums.PriceHistorySinglePostPageNumber)
	case "btn_bookmark":
	}

	return StaticBtns.HandleIfClickBackBtn(btnKey, PostEnums.MainPostSelectedReportPageNumber)
}

var _ Page.PageInterface = &ShowSinglePostPage{}
