package Range

import (
	Models2 "project-root/modules/report/DB/Models"
	"project-root/modules/report/Enums"
	"project-root/modules/report/Http/Controllers/Filter"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
)

type ApartmentUpdateFilterPage struct{}

func (p *ApartmentUpdateFilterPage) PageNumber() int {
	return Enums.ApartmentUpdateFilterPageNumber
}

func (p *ApartmentUpdateFilterPage) GeneratePage(telSession *Models.TelSession) *Page.PageContentOV {
	return Filter.GetGeneratePageToggleUpdateFilter("Apartment")
}

func (p *ApartmentUpdateFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *ApartmentUpdateFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	return Filter.OnClickInlineBtnToggleUpdateFilter(btnKey, telSession, func(isActive *int, filter *Models2.ReportFilter) {
		filter.IsApartment = isActive
	})
}

var _ Page.PageInterface = &ApartmentUpdateFilterPage{}
