package Range

import (
	tele "gopkg.in/telebot.v4"
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type StorageUpdateFilterPage struct{}

func (p *StorageUpdateFilterPage) PageNumber() int {
	return Enums.StorageUpdateFilterPageNumber
}

func (p *StorageUpdateFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	return Filter.GetGeneratePageToggleUpdateFilter("Storage")
}

func (p *StorageUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *StorageUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnClickInlineBtnToggleUpdateFilter(btnKey, telSession, func(isActive *int, filter *Models2.ReportFilter) {
		filter.Storage = isActive
	})
}

var _ Page.PageInterface = &StorageUpdateFilterPage{}
