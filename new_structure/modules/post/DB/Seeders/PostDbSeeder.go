package Seeders

import (
	"gorm.io/gorm"
	"project-root/modules/post/DB/Models"
	PostEnums "project-root/modules/post/Enums"
	"project-root/modules/post/Facades"
	"project-root/modules/report/Enums"
	"project-root/sys-modules/database/Lib"
	"time"
)

type PostDbSeeder struct{}

func (s PostDbSeeder) Name() string {
	return "300_posts_table"
}

func (s PostDbSeeder) Handle(db *gorm.DB) {
	Facades.PostRepo().Truncate()
	now := time.Now()
	//sourceSite, err := Facades2.SourceSiteRepo().FindByID(1)
	//if err != nil {
	//	panic("source site id not found")
	//}
	posts := []Models.Post{
		{
			SrcSitesID:       1,
			SrcSite:          nil,
			UsersID:          nil,
			User:             nil,
			Status:           PostEnums.ProcessingPostStatus,
			ExternalSiteID:   ptrString("wZr4S2yn"),
			Title:            "۹۸متر آپارتمان فول ۲پارکینگ م رسالت خ نیرودریایی",
			Description:      ptrString("توضیحات\n\n۹۸متر آپارتمان با ۲پارکینگ و آسانسور و انباری و بالکن\nبدون مالک \nتخلیه\n آماده تحویل \nخوش نقشه و شیک تمیز \nمشاعات راه پله ها عالی \nغرق نور طبقه ۴\nلوکیشن عالی خلوت و دنج\nمیدان رسالت خیابان نیرودریایی کوچه ارجمندی نزدیک پارک دنج و خلوت"),
			Price:            ptrInt64(2500000000),
			PriceHistory:     map[string]interface{}{"2023-11-01 13:12:07": "2400000000", "2023-11-10 17:22:00": "2500000000"},
			MainIMG:          ptrString("https://s100.divarcdn.com/static/photo/neda/post/-Ms9Mk1LQf9ELOnkcoU4oA/8c62bb3b-4e60-4d69-b696-9f4597d3aad2.jpg"),
			GalleryIMGs:      nil, //map[string]interface{}{"img1": "https://example.com/gallery1.jpg", "img2": "https://example.com/gallery2.jpg"},
			SellerName:       nil,
			LandArea:         ptrFloat64(120.5),
			BuiltYear:        ptrInt(2015),
			Rooms:            ptrInt(3),
			IsApartment:      ptrBool(true),
			DealType:         ptrInt(Enums.RentDealType),
			Floors:           ptrInt(10),
			Elevator:         ptrBool(true),
			Storage:          ptrBool(false),
			Location:         nil,
			PostDate:         ptrTime(now.Add(-48 * time.Hour)), // 2 روز پیش
			CityName:         ptrString("تهران"),
			NeighborhoodName: ptrString("ونک"),
			DeletedAt:        nil,
			CreatedAt:        now.Add(-72 * time.Hour),         // 3 روز پیش
			UpdateAt:         ptrTime(now.Add(-2 * time.Hour)), // 2 ساعت پیش
		},
		{
			SrcSitesID:       1,
			SrcSite:          nil,
			UsersID:          nil,
			User:             nil,
			Status:           PostEnums.SuccessfulPostStatus, // successful
			ExternalSiteID:   ptrString("wZ4I1C-i"),
			Title:            "۸۵متر دو خواب تکواحد طبقه ۴ با اسانسور",
			Description:      ptrString("توضیحات\n\n۸۵ متر دو خواب \n\nتکواحد \n\nطبقه ۴ با اسانسور \n\nواحد جنوبی \n\nسند به صورت قولنامه ای \n\nفروشنده واقعی \n\nمشاور فرهاد بیات"),
			Price:            ptrInt64(3500000000),
			PriceHistory:     map[string]interface{}{"2023-11-05 20:00:00": "3400000000", "2023-11-15 16:22:00": "3500000000"},
			MainIMG:          nil,
			GalleryIMGs:      nil,
			SellerName:       nil,
			LandArea:         ptrFloat64(400),
			BuiltYear:        ptrInt(2010),
			Rooms:            ptrInt(5),
			IsApartment:      ptrBool(false),
			DealType:         ptrInt(Enums.RentDealType), // اجاره
			Floors:           ptrInt(2),
			Elevator:         ptrBool(false),
			Storage:          ptrBool(true),
			Location:         nil,
			PostDate:         ptrTime(now.Add(-24 * time.Hour)), // 1 روز پیش
			CityName:         ptrString("تهران"),
			NeighborhoodName: ptrString("چهار باغ"),
			DeletedAt:        nil,
			CreatedAt:        now.Add(-48 * time.Hour),         // 2 روز پیش
			UpdateAt:         ptrTime(now.Add(-1 * time.Hour)), // 1 ساعت پیش
		},
	}

	for _, post := range posts {
		Facades.PostRepo().Create(&post)
	}
}

func ptrString(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

func ptrInt64(i int64) *int64 {
	return &i
}

func ptrFloat64(f float64) *float64 {
	return &f
}

func ptrBool(b bool) *bool {
	return &b
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

var _ Lib.DbSeederInterface = &PostDbSeeder{}
