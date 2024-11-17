package user

import (
	"project-root/app"
	"project-root/modules/report/DB/Migrations"
	"project-root/modules/report/Http/Controllers"
	"project-root/modules/report/Http/Controllers/Create"
	Global "project-root/modules/report/Http/Controllers/Filter/Global"
	"project-root/modules/report/Http/Controllers/Filter/Range"
	Toggle "project-root/modules/report/Http/Controllers/Filter/Toggle"
	Update "project-root/modules/report/Http/Controllers/Update"
	SysDatabase "project-root/sys-modules/database/Lib"
	"project-root/sys-modules/telebot/Lib/Page"
)

func (s *ReportServiceProvider) Register() {
	// Migrations
	SysDatabase.RegisterMigration(&Migrations.CreateReportTable{})
	SysDatabase.RegisterMigration(&Migrations.CreateReportFilterTable{})

	// Pages
	Page.RegisterPages([]Page.PageInterface{
		&Controllers.MainReportUserPage{},

		// create
		&Create.TitleCreateReportPage{},
		&Create.IsNotificationCreateReportPage{},

		// update
		&Controllers.MainSelectedReportPage{},
		&Update.TitleUpdateReportPage{},
		&Update.IsNotificationUpdateReportPage{},

		// filter
		&Controllers.MainUpdateReportFilterPage{},
		&Range.RoomUpdateFilterPage{},
		&Range.PriceUpdateFilterPage{},
		&Range.AreaUpdateFilterPage{},
		&Range.BuiltUpdateFilterPage{},
		&Range.FloorUpdateFilterPage{},
		&Toggle.StorageUpdateFilterPage{},
		&Toggle.ElevatorUpdateFilterPage{},
		&Toggle.ApartmentUpdateFilterPage{},
		&Global.CityUpdateFilterPage{},
		&Global.NeighborhoodUpdateFilterPage{},
		&Global.DealTypeUpdateFilterPage{},
		&Global.PublishDateUpdateFilterPage{},
	})
}

func (s *ReportServiceProvider) Boot() {

}

type ReportServiceProvider struct{}

var _ app.ServiceProviderInterface = &ReportServiceProvider{}
