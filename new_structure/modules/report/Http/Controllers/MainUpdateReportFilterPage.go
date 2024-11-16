package Controllers

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	ReportModels "project-root/modules/report/DB/Models"
	ReportEnums "project-root/modules/report/Enums"
	"project-root/modules/report/Facades"
	"project-root/modules/user/DB/Models"
	"project-root/sys-modules/telebot/Lib/Page"
	"time"
)

type MainUpdateReportFilterPage struct{}

func (p *MainUpdateReportFilterPage) PageNumber() int {
	return ReportEnums.MainUpdateReportFilterPageNumber
}

func (p *MainUpdateReportFilterPage) GeneratePage(telSession *Models.TelSession) (string, *tele.ReplyMarkup) {
	var newReplyMarkup = &tele.ReplyMarkup{}
	btnRoom := newReplyMarkup.Data("Room", "btn_room")
	btnLandArea := newReplyMarkup.Data("Land Area", "btn_land_area")
	btnPrice := newReplyMarkup.Data("Price", "btn_price")

	btnPublishDate := newReplyMarkup.Data("Publish Date", "btn_publish_date")
	btnStorage := newReplyMarkup.Data("Storage", "btn_storage")
	btnFloor := newReplyMarkup.Data("Floor", "btn_floor")

	btnApartment := newReplyMarkup.Data("Apartment", "btn_apartment")
	btnBuiltYear := newReplyMarkup.Data("Built year", "btn_built_year")
	btnElevator := newReplyMarkup.Data("Elevator", "btn_elevator")

	btnSourceSites := newReplyMarkup.Data("Source Sites", "btn_source_sites")
	btnDealType := newReplyMarkup.Data("Deal Type", "btn_deal_type")

	btnCity := newReplyMarkup.Data("City", "btn_city")
	btnNeighborhood := newReplyMarkup.Data("Neighborhood", "btn_neighborhood")

	btnBack := newReplyMarkup.Data("Back", "btn_back")

	newReplyMarkup.Inline(
		newReplyMarkup.Row(btnRoom, btnLandArea, btnPrice),
		newReplyMarkup.Row(btnPublishDate, btnStorage, btnFloor),
		newReplyMarkup.Row(btnApartment, btnBuiltYear, btnElevator),
		newReplyMarkup.Row(btnSourceSites, btnDealType),
		newReplyMarkup.Row(btnCity, btnNeighborhood),
		newReplyMarkup.Row(btnBack),
	)

	filter, _ := Facades.ReportFilterRepo().FindByID(telSession.GetReportTempData().FilterId)
	return getFilterBody(filter), newReplyMarkup
}

func getFilterBody(filter *ReportModels.ReportFilter) string {
	report := "Making the filter - Click on each filter button to create or change a filter :\n"

	priceRange := fmt.Sprintf("Price range : %s - %s\n", getString(filter.PriceMin), getString(filter.PriceMax))
	landAreaRange := fmt.Sprintf("Land area range : %s - %s\n", getString(filter.AreaMin), getString(filter.AreaMax))
	roomRange := fmt.Sprintf("Room range : %s - %s\n", getString(filter.RoomCountMin), getString(filter.RoomCountMax))
	floorRange := fmt.Sprintf("Floor range : %s - %s\n", getString(filter.FloorCountMin), getString(filter.FloorCountMax))
	haveStorage := fmt.Sprintf("have storage : %s\n", getString(filter.Storage))
	publishDataRange := fmt.Sprintf("Publish data range : %s - %s\n", getString(filter.PostStartDate), getString(filter.PostEndDate))
	haveElevator := fmt.Sprintf("have elevator : %s\n", getString(filter.Elevator))
	builtYearRange := fmt.Sprintf("Built year range : %s - %s\n", getString(filter.BuiltStart), getString(filter.BuiltEnd))
	isApartment := fmt.Sprintf("Is apartment : %s\n", getString(filter.IsApartment))
	cty := fmt.Sprintf("City name : %s\n", getString(filter.CityName))
	neighborhood := fmt.Sprintf("Neighborhood name : %s\n", getString(filter.NeighborhoodName))
	delType := fmt.Sprintf("Deal type : %s\n", getString(filter.DealType))
	sourcSite := "Source site : divar\n"

	body := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
		report, priceRange, landAreaRange, roomRange, floorRange,
		haveStorage, publishDataRange, haveElevator,
		builtYearRange, isApartment, cty, neighborhood, delType, sourcSite)

	return body
}

func getString(value interface{}) string {
	switch v := value.(type) {
	case *int:
		if v == nil {
			return ""
		}
		return fmt.Sprintf("%d", *v)
	case *string:
		if v == nil {
			return ""
		}
		return *v
	case *time.Time:
		if v == nil {
			return ""
		}
		return v.Format("2006-01-02 15:04:05")
	case *bool:
		if v == nil {
			return ""
		}
		return fmt.Sprintf("%t", *v)
	default:
		return ""
	}
}

func (p *MainUpdateReportFilterPage) OnInput(value string, telSession *Models.TelSession) Page.PageInterface {
	return nil
}

func (p *MainUpdateReportFilterPage) OnClickInlineBtn(btnKey string, telSession *Models.TelSession) Page.PageInterface {
	switch btnKey {
	case "btn_room":
		return Page.GetPage(ReportEnums.RoomUpdateFilterPageNumber)
	case "btn_land_area":
		return Page.GetPage(ReportEnums.AreaUpdateFilterPageNumber)
	case "btn_price":
		return Page.GetPage(ReportEnums.PriceUpdateFilterPageNumber)
	case "btn_floor":
		return Page.GetPage(ReportEnums.FloorUpdateFilterPageNumber)
	case "btn_built_year":
		return Page.GetPage(ReportEnums.BuiltUpdateFilterPageNumber)
	case "btn_storage":
		return Page.GetPage(ReportEnums.StorageUpdateFilterPageNumber)
	case "btn_apartment":
		return Page.GetPage(ReportEnums.ApartmentUpdateFilterPageNumber)
	case "btn_elevator":
		return Page.GetPage(ReportEnums.ElevatorUpdateFilterPageNumber)
	case "btn_city":
		return Page.GetPage(ReportEnums.CityUpdateFilterPageNumber)
	case "btn_neighborhood":
		return Page.GetPage(ReportEnums.NeighborhoodUpdateFilterPageNumber)
	case "btn_deal_type":
		return Page.GetPage(ReportEnums.DealTypeUpdateFilterPageNumber)
	case "btn_source_sites":
		telSession.GetGeneralTempData().LastMessage = "We only support Divar!"
		return Page.GetPage(ReportEnums.MainUpdateReportFilterPageNumber)
	case "btn_publish_date":
		return Page.GetPage(ReportEnums.PublishDateUpdateFilterPageNumber)
	case "btn_back":
		return Page.GetPage(ReportEnums.MainReportUserPageNumber)
	}
	return nil
}

var _ Page.PageInterface = &MainUpdateReportFilterPage{}
