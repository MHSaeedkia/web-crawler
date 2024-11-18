package Controllers

import (
	"fmt"
	PostModels "project-root/modules/post/DB/Models"
	PostEnums "project-root/modules/post/Enums"
	Facades2 "project-root/modules/post/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"project-root/sys-modules/telebot/Lib/StaticBtns"
	"project-root/sys-modules/time/Lib"
	"sort"
	"time"
)

type PriceHistorySinglePostPage struct{}

func (p *PriceHistorySinglePostPage) PageNumber() int {
	return PostEnums.PriceHistorySinglePostPageNumber
}

func (p *PriceHistorySinglePostPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	post, _ := Facades2.PostRepo().FindByID(telSession.GetPostTempData().PostId)
	fmt.Println(GeneratePriceHistory(*post))
	return &Page.PageContentOV{
		Message:     GeneratePriceHistory(*post),
		ReplyMarkup: StaticBtns.GetBackStaticBtn(),
	}
}

func GeneratePriceHistory(post PostModels.Post) string {
	var output string
	output = "Price History -->\n"
	output += post.Title + "\n\n"
	keys := make([]string, 0, len(post.PriceHistory))
	for key := range post.PriceHistory {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", key)
		strDate := "The date format is not valid"
		if err == nil {
			strDate = Lib.FormatTimeAgo(parsedTime)
		}

		price := post.PriceHistory[key]
		output += fmt.Sprintf("%s\n%v\n---------------\n", strDate, price)
	}

	return output
}

func (p *PriceHistorySinglePostPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *PriceHistorySinglePostPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return StaticBtns.HandleIfClickBackBtn(btnKey, PostEnums.ShowSinglePostPageNumber)
}

var _ Page.PageInterface = &PriceHistorySinglePostPage{}
