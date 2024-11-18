package StaticBtns

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"project-root/sys-modules/telebot/Lib/Page"
	"regexp"
)

/*type PaginationPage struct {
	ReplyMarkupData PaginationReplyMarkupData
	OnClickBtnData  PaginationOnClickBtnData
}*/

// ------------ ReplyMarkupData
type PaginationReplyMarkupData struct {
	Items        []PaginationReplyMarkupItem
	StaticRowBtn []tele.Row
}

type PaginationReplyMarkupItem struct {
	ID    int
	Title string
}

// ------------ OnClickBtn
type PaginationOnClickBtnData struct {
	PrefixBtnKey       string
	CurrentPageNumber  int
	BackPageNumber     int
	GetPageNumberSaved func() int
	SavePageNumber     func(pageNum int)
	OnSelectItemId     func(itemId int) Page.PageInterface
}

func (pagination *PaginationReplyMarkupData) GetReplyMarkup(prefixBtnKey string) *tele.ReplyMarkup {
	var newReplyMarkup = &tele.ReplyMarkup{}

	var rows []tele.Row

	// dynamic btn
	for _, item := range pagination.Items {
		btn := newReplyMarkup.Data(item.Title, fmt.Sprintf("btn_show_%s_%d", prefixBtnKey, item.ID))
		rows = append(rows, newReplyMarkup.Row(btn))
	}

	// static pagination btn
	btnPreviousPage := newReplyMarkup.Data("previous page", "btn_previous_page")
	btnNextPage := newReplyMarkup.Data("next page", "btn_next_page")
	btnBack := newReplyMarkup.Data("Back", "btn_back")
	rows = append(rows, newReplyMarkup.Row(btnPreviousPage, btnNextPage))

	// static custom btn
	for _, row := range pagination.StaticRowBtn {
		rows = append(rows, row)
	}

	// --
	rows = append(rows, newReplyMarkup.Row(btnBack))

	newReplyMarkup.Inline(rows...)
	return newReplyMarkup
}

func (pagination *PaginationOnClickBtnData) HandleInputPagination(btnKey string) Page.PageInterface {
	switch btnKey {
	case "btn_next_page":
		pagination.SavePageNumber(pagination.GetPageNumberSaved() + 1)
		return Page.GetPage(pagination.CurrentPageNumber)
	case "btn_previous_page":
		if pagination.GetPageNumberSaved() > 1 {
			pagination.SavePageNumber(pagination.GetPageNumberSaved() - 1)
		} else {
			pagination.SavePageNumber(1)
		}
		return Page.GetPage(pagination.CurrentPageNumber)
	case "btn_back":
		return Page.GetPage(pagination.BackPageNumber)
	default:
		// dynamic btn
		itemId, err := parseBtnKey(btnKey, "post")
		if err == nil {
			return pagination.OnSelectItemId(itemId)
		}
	}
	return nil
}

func parseBtnKey(btnKey, prefix string) (int, error) {
	pattern := fmt.Sprintf(`^btn_show_%s_(\d+)`, prefix)
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(btnKey)

	if len(matches) > 0 {
		var itemIdID int
		fmt.Sscanf(matches[1], "%d", &itemIdID)
		return itemIdID, nil
	}
	return 0, fmt.Errorf("input does not match format")
}
