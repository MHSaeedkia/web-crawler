package telegrambot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

const (
	USER       = "User"
	ADMIN      = "Admin"
	SUPERADMIN = "SuperAdmin"
)

const (
	rent = 0
	buy  = 1
	sell = 2
)

const (
	divar    = 0
	sheypoor = 1
)

type user struct {
	userName string
	password string
	userType string
	email    string
}

type filters struct {
	priceRange       []int
	landAreaRange    []int
	roomRange        []int
	floorRange       []int
	haveStorage      bool
	publishDataRange []string //[]time.Time
	haveElevator     bool
	builtYearRange   []int
	isApartment      bool
	city             string
	dealType         int
	sourceSite       []int
}

var (
	bot *telebot.Bot

	menu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	selector = &telebot.ReplyMarkup{RemoveKeyboard: true}

	data    = make(map[int64]user)
	message = make(map[int64]*telebot.Message)
	filter  = make(map[int64]filters)

	// inline btn
	register = selector.Data("Register", "Register")
	login    = selector.Data("Login", "Login")

	profile            = selector.Data("Profile", "Profile")
	bookmark           = selector.Data("Bookmark", "Bookmark")
	prevPageMyBkMark   = selector.Data("Previuse page", "Previuse_page_my_bookmark")
	nexPageMyBkMark    = selector.Data("Next page", "Next_page_my_bookmark")
	prevPageGlobBkMark = selector.Data("Previuse page", "Previuse_page_global_bookmark")
	nexPageGlobBkMark  = selector.Data("Next page", "Next_page_global_bookmark")

	monitoring      = selector.Data("Monitoring", "Monitoring")
	prevPageMonitor = selector.Data("Previuse page", "Previuse_page_monitoring")
	nexPageMonitor  = selector.Data("Next page", "Next_page_monitoring")

	post               = selector.Data("Post", "Post")
	prevPagePost       = selector.Data("Previuse page", "Previuse_page_post")
	nexPagePost        = selector.Data("Next page", "Next_page_post")
	selectPost         = selector.Data("", "Select_Post")
	prevSelectPagePost = selector.Data("Previuse page", "Previuse_Select_Page_Post")
	nexSelectPagePost  = selector.Data("Next page", "Next_Select_Page_Post")
	newPost            = selector.Data("Create New Post", "Create_New_Post")
	backPost           = selector.Data("Back", "Back_Post")

	export            = selector.Data("Export", "Export")
	report            = selector.Data("Report", "Report")
	newReport         = selector.Data("Create New Report", "Create_New_Report")
	newReportBack     = selector.Data("Back", "Back_New_report")
	newReportNo       = selector.Data("No", "New_report_No")
	newReportYes      = selector.Data("Yes", "New_report_Yes")
	newReportNotifNo  = selector.Data("No", "New_report_Notif_No")
	newReportNotifYes = selector.Data("Yes", "New_report_Notif_Yes")
	prevPageReport    = selector.Data("Previuse page", "Previuse_page_report")
	nexPageReport     = selector.Data("Next page", "Next_page_report")
	room              = selector.Data("Room", "Room")
	landArea          = selector.Data("Land Area", "Land_Area")
	price             = selector.Data("Price", "Price")
	publishData       = selector.Data("Publish Data", "Publish_Data")
	storage           = selector.Data("Storage", "Storage")
	floor             = selector.Data("Floor", "Floor")
	apartment         = selector.Data("Apartment", "Apartment")
	builtYear         = selector.Data("Built Year", "Built_Year")
	elevator          = selector.Data("Elevator", "Elevator")
	sourceSite        = selector.Data("Source Site", "Source_Sites")
	dealType          = selector.Data("Deal Type", "Deal_Type")
	city              = selector.Data("City", "City")
	finish            = selector.Data("Finish", "Finish")
	roomRemove        = selector.Data("Remove", "Remove_Room")
	landAreaRemove    = selector.Data("Remove", "Remove_Land_Area")
	priceRemove       = selector.Data("Remove", "Remove_Price")
	publishDataRemove = selector.Data("Remove", "Remove_Publish_Data")
	storageRemove     = selector.Data("Remove", "Remove_Storage")
	floorRemove       = selector.Data("Remove", "Remove_Floor")
	apartmentRemove   = selector.Data("Remove", "Remove_Apartment")
	builtYearRemove   = selector.Data("Remove", "Remove_Built_Year")
	elevatorRemove    = selector.Data("Remove", "Remove_Elevator")
	sourceSiteRemove  = selector.Data("Remove", "Remove_Source_Sites")
	dealTypeRemove    = selector.Data("Remove", "Remove_Deal_Type")
	cityRemove        = selector.Data("Remove", "Remove_City")
	backFilter        = selector.Data("Back", "Back_Filter")
	backFilterEdit    = selector.Data("Back", "Back_Filter_Edit")
	selectReport      = selector.Data("", "Select_Report")
	editReport        = selector.Data("Edit", "Edit_Report")
	deleteReport      = selector.Data("Delete", "Delete_Report")

	logout = selector.Data("Logout", "Logout")

	monitoringAdmin      = selector.Data("Monitoring", "MonitoringAdmin")
	usr                  = selector.Data("User", "User")
	accessToUserAccount  = selector.Data("Access To User Account", "AccessToUserAccount")
	prevPageMonitorAdmin = selector.Data("Previuse page", "Previuse_Page_Monitoring_Admin")
	nexPageMonitorAdmin  = selector.Data("Next page", "Next_Page_Monitoring_Admin")
	refreshMonitorAdmin  = selector.Data("Refresh", "Refresh_Monitor_Admin")
	backMonitorAdmin     = selector.Data("Back", "Back_Monitor_Admin")
	prevUser             = selector.Data("Previuse page", "Previuse_Page_User")
	nexUser              = selector.Data("Next page", "Next_Page_User")

	monitoringSuperAdmin      = selector.Data("Monitoring", "MonitoringSuperAdmin")
	prevPageMonitorSuperAdmin = selector.Data("Previuse page", "Previuse_Page_Monitoring_Super_Admin")
	nexPageMonitorSuperAdmin  = selector.Data("Next page", "Next_Page_Monitoring_Super_Admin")
	refreshMonitorSuperAdmin  = selector.Data("Refresh", "Refresh_Monitor_Super_Admin")
	backMonitorSuperAdmin     = selector.Data("Back", "Back_Monitor_Super_Admin")
	setting                   = selector.Data("Setting", "Setitng")
	admin                     = selector.Data("Admin", "Admin")
	crawlerInterval           = selector.Data("Crawler inteval", "Crawler_Interval")
	maxPosts                  = selector.Data("Max Posts", "Max_Posts")
	prevPageAdminSuperAdmin   = selector.Data("Previuse page", "Previuse_Page_Admin_Super_Admin")
	nexPageAdminSuperAdmin    = selector.Data("Next page", "Next_Page_Admin_Super_Admin")
	createNewAdmin            = selector.Data("Create New Admin", "Create_New_Admin")

	globalBookmark = selector.Data("Global Bookmarks", "Global_Bookmarks")
	myBookmark     = selector.Data("My Bookmarks", "My_Bookmarks")
	back           = selector.Data("Back", "Back")
	refresh        = selector.Data("Refresh", "Refresh")
	skip           = selector.Data("Skip", "Skip")

	// reply btn
	cancle = menu.Text("Cancle")

	// state
	loginStatus            = false
	registerUserStatus     = false
	registerAdminStatus    = false
	newReportStatus        = false
	priceRangeStatus       = false
	landAreaRangeStatus    = false
	roomRangeStatus        = false
	floorRangeStatus       = false
	haveStorageStatus      = false
	publishDataRangeStatus = false
	haveElevatorStatus     = false
	builtYearRangeStatus   = false
	isApartmentStatus      = false
	cityStatus             = false
	dealTypeStatus         = false
	sourceSiteStatus       = false

	newPostStatus = false
)

func initServer() (*telebot.Bot, error) {
	err := godotenv.Load("/tmp/.env")
	if err != nil {
		return nil, err
	}

	pref := telebot.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err = telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return bot, nil
}

func StartBot() {
	bot, err := initServer()
	if err != nil {
		panic(err)
	}

	// endpoints .
	bot.Handle("/start", start)

	bot.Handle(&register, registerBtn)
	bot.Handle(&login, loginBtn)

	bot.Handle(&monitoring, monitoringBtn)
	bot.Handle(&bookmark, globalBookmarkBtn)
	bot.Handle(&post, postBtn)
	bot.Handle(&selectPost, selectPostBtn)
	bot.Handle(&newPost, newPostBtn)
	bot.Handle(&backPost, backPostBtn)

	bot.Handle(&report, reportBtn)
	bot.Handle(&newReport, newReportBtn)
	bot.Handle(&newReportBack, newReportBtn)
	bot.Handle(&newReportNo, newReportNoBtn)
	bot.Handle(&newReportYes, newReportYesBtn)
	bot.Handle(&newReportNotifNo, newReporNotifNoBtn)
	bot.Handle(&newReportNotifYes, newReporNotifYesBtn)
	bot.Handle(&room, roomBtn)
	bot.Handle(&landArea, landAreaBtn)
	bot.Handle(&price, priceBtn)
	bot.Handle(&publishData, publishDataBtn)
	bot.Handle(&storage, storageBtn)
	bot.Handle(&floor, floorBtn)
	bot.Handle(&apartment, apartmentBtn)
	bot.Handle(&builtYear, builtYearBtn)
	bot.Handle(&elevator, elevatorBtn)
	bot.Handle(&sourceSite, sourceSiteBtn)
	bot.Handle(&dealType, dealTypeBtn)
	bot.Handle(&city, cityBtn)
	bot.Handle(&finish, finishBtn)
	bot.Handle(&backFilter, backFilterBtn)
	bot.Handle(&roomRemove, roomRemoveBtn)
	bot.Handle(&landAreaRemove, landAreaRemoveBtn)
	bot.Handle(&priceRemove, priceRemoveBtn)
	bot.Handle(&publishDataRemove, publishDataRemoveBtn)
	bot.Handle(&storageRemove, storageRemoveBtn)
	bot.Handle(&floorRemove, floorRemoveBtn)
	bot.Handle(&apartmentRemove, apartmentRemoveBtn)
	bot.Handle(&builtYearRemove, builtYearRemoveBtn)
	bot.Handle(&elevatorRemove, elevatorRemoveBtn)
	bot.Handle(&sourceSiteRemove, sourceSiteRemoveBtn)
	bot.Handle(&dealTypeRemove, dealTypeRemoveBtn)
	bot.Handle(&cityRemove, cityRemoveBtn)
	bot.Handle(&backFilterEdit, backFilterEditBtn)
	bot.Handle(&selectReport, selectReportBtn)
	bot.Handle(&deleteReport, deleteReportBtn)
	bot.Handle(&editReport, editReportBtn)

	bot.Handle(&monitoringAdmin, monitoringAdminBtn)
	bot.Handle(&monitoringSuperAdmin, monitoringSuperAdminBtn)
	bot.Handle(&usr, usrBtn)
	bot.Handle(&setting, settingBtn)
	bot.Handle(&admin, adminBtn)
	bot.Handle(&accessToUserAccount, accessToUserAccountBtn)

	bot.Handle(&back, backBtn)
	bot.Handle(&myBookmark, myBookmarkBtn)
	bot.Handle(&globalBookmark, globalBookmarkBtn)
	bot.Handle(&newPost, newPostBtn)

	bot.Handle(telebot.OnText, onText)

	bot.Handle(&profile, profileBtn)

	// start bot
	bot.Start()
}

func start(c telebot.Context) error {
	selector.Inline(selector.Row(register, login))
	return c.Send(c.Sender(), "Select an operation :", selector)
}

func registerBtn(c telebot.Context) error {
	c.Delete()
	registerUserStatus = true
	userId := c.Sender().ID
	data[userId] = user{}

	resp, err := bot.Send(
		c.Sender(),
		"Please enter your username :",
	)
	if err != nil {
		return err
	}
	message[userId] = resp
	return nil
}

func loginBtn(c telebot.Context) error {
	c.Delete()
	loginStatus = true
	userId := c.Sender().ID
	data[userId] = user{}

	resp, err := bot.Send(
		c.Sender(),
		"Please enter your username :",
	)
	if err != nil {
		return err
	}
	message[userId] = resp
	return nil
}

func logoutBtn(c telebot.Context) error {
	selector.Inline(selector.Row(register, login))
	return c.Edit("Select an operation :", selector)
}

func monitoringAdminBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	log := fmt.Sprintf("crawler logs activities : \n%s", "log ..")
	selector.Inline(
		selector.Row(prevPageMonitorAdmin, page, nexPageMonitorAdmin),
		selector.Row(refreshMonitorAdmin),
		selector.Row(backMonitorAdmin),
	)

	return c.Edit(
		log,
		selector,
	)
}

func accessToUserAccountBtn(c telebot.Context) error {
	log := "Enter the user's username to loginto their account without a password : "
	selector.Inline(
		selector.Row(backMonitorAdmin),
	)

	return c.Edit(
		log,
		selector,
	)
}

func usrBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	log := fmt.Sprintf("To change the status of each admin , click on the botton in front of it.")

	// copule of inline btn .

	selector.Inline(
		selector.Row(prevUser, page, nexUser),
		selector.Row(backMonitorAdmin),
	)

	return c.Edit(
		log,
		selector,
	)
}

func monitoringSuperAdminBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	log := fmt.Sprintf("crawler logs activities : \n%s", "log ..")
	selector.Inline(
		selector.Row(prevPageMonitorSuperAdmin, page, nexPageMonitorSuperAdmin),
		selector.Row(refreshMonitorSuperAdmin),
		selector.Row(backMonitorSuperAdmin),
	)

	return c.Edit(
		log,
		selector,
	)
}

func settingBtn(c telebot.Context) error {
	setting := fmt.Sprintf("To set each value , click on button  : \n\ncrawler interval : %v min \nmax posts : %v\n", 250, 20000)
	selector.Inline(
		selector.Row(crawlerInterval, maxPosts),
		selector.Row(backMonitorSuperAdmin),
	)

	return c.Edit(
		setting,
		selector,
	)
}

func adminBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	admin := fmt.Sprintf("To change the status of each admin , click on the button in front of it : ")

	// some inlit btn .

	selector.Inline(
		selector.Row(prevPageAdminSuperAdmin, page, nexPageAdminSuperAdmin),
		selector.Row(createNewAdmin),
		selector.Row(backMonitorSuperAdmin),
	)

	return c.Edit(
		admin,
		selector,
	)
}

func createNewAdminBtn(c telebot.Context) error {
	c.Delete()
	registerAdminStatus = true
	userId := c.Sender().ID
	data[userId] = user{}

	resp, err := bot.Send(
		c.Sender(),
		"Please enter your username :",
	)
	if err != nil {
		return err
	}
	message[userId] = resp
	return nil
}

func monitoringBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	log := fmt.Sprintf("crawler logs activities : \n%s", "log ..")
	selector.Inline(
		selector.Row(prevPageMonitor, page, nexPageMonitor),
		selector.Row(refresh),
		selector.Row(back),
	)

	return c.Edit(
		log,
		selector,
	)
}

func myBookmarkBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	bookMarks := fmt.Sprintf("My bookmark  : \n%s", "bookmark ..")
	selector.Inline(
		selector.Row(prevPageMyBkMark, page, nexPageMyBkMark),
		selector.Row(globalBookmark),
		selector.Row(back),
	)

	return c.Edit(
		bookMarks,
		selector,
	)
}

func globalBookmarkBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	bookMarks := fmt.Sprintf("global bookmark : \n%s", "bookmark ..")
	selector.Inline(
		selector.Row(prevPageGlobBkMark, page, nexPageGlobBkMark),
		selector.Row(myBookmark),
		selector.Row(back),
	)

	return c.Edit(
		bookMarks,
		selector,
	)
}

func postBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	posts := "Select a repost to display its posts : "

	// copule of inline btn .

	selector.Inline(
		selector.Row(prevPagePost, page, nexPagePost),
		selector.Row(back),
	)

	return c.Edit(
		posts,
		selector,
	)
}

func selectPostBtn(c telebot.Context) error {
	message := fmt.Sprintf("The report \"%s\" has been selected , select a post to view , edit or delete :", "vespa 200m")

	// some inline btn .
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	selector.Inline(
		selector.Row(prevSelectPagePost, page, nexSelectPagePost),
		selector.Row(newPost),
		selector.Row(backPost),
	)
	return c.Edit(
		message,
		selector,
	)

}

func newPostBtn(c telebot.Context) error {
	newPostStatus = true
	posts := "Enter a name for your post : "

	// copule of inline btn .

	selector.Inline(
		selector.Row(newPost),
		selector.Row(backPost),
	)

	return c.Edit(
		posts,
		selector,
	)
}

func reportBtn(c telebot.Context) error {
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	report := fmt.Sprintf("The report \"%s\" has been selected , select a post to view , edit or delere :", "vespa 200 divar and sheypoor")

	// copule of inline btn .

	selector.Inline(
		selector.Row(prevPageReport, page, nexPageReport),
		selector.Row(newReport),
		selector.Row(back),
	)

	return c.Edit(
		report,
		selector,
	)
}

func newReportBtn(c telebot.Context) error {
	c.Delete()
	newReportStatus = true
	userId := c.Sender().ID
	data[userId] = user{}

	selector.Inline(
		selector.Row(back),
	)

	resp, err := bot.Send(
		c.Sender(),
		"Type a name for your report , this name should be unique :",
		selector,
	)
	if err != nil {
		return err
	}
	message[userId] = resp
	return nil
}

func newReportNoBtn(c telebot.Context) error {
	// do something
	loginMessage := "If this is checked again (watch list) , should the notification be sent to robot ?"
	selector.Inline(
		selector.Row(newReportNotifNo, newReportNotifYes),
		selector.Row(newReportBack),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func newReportYesBtn(c telebot.Context) error {
	// do something
	loginMessage := "If this is checked again (watch list) , should the notification be sent to robot ?"
	selector.Inline(
		selector.Row(newReportNotifNo, newReportNotifYes),
		selector.Row(newReportBack),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func newReporNotifNoBtn(c telebot.Context) error {
	// do something
	return reportFilter(c)
}

func newReporNotifYesBtn(c telebot.Context) error {
	// do something
	return reportFilter(c)
}

func reportFilter(c telebot.Context) error {
	userId := c.Sender().ID
	fltr, exist := filter[userId]
	if !exist {
		filter[userId] = filters{}
	}
	report := fmt.Sprintf("Making the filter - Click on each filter button to create or change a filter :\n")
	priceRange := fmt.Sprintf("Price range : %v - %v\n", fltr.priceRange[0], fltr.priceRange[1])
	landAreaRange := fmt.Sprintf("Land area range : %v - %v\n", fltr.landAreaRange[0], fltr.landAreaRange[1])
	roomRange := fmt.Sprintf("Room range : %v - %v\n", fltr.roomRange[0], fltr.roomRange[1])
	floorRange := fmt.Sprintf("Floor range : %v - %v\n", fltr.floorRange[0], fltr.floorRange[1])
	haveStorage := fmt.Sprintf("have storage : %t\n", fltr.haveStorage)
	publishDataRange := fmt.Sprintf("Publish data range : %v - %v\n", fltr.publishDataRange[0], fltr.publishDataRange[1])
	haveElevator := fmt.Sprintf("have elevator : %t\n", fltr.haveElevator)
	builtYearRange := fmt.Sprintf("Built year range : %v - %v\n", fltr.builtYearRange[0], fltr.builtYearRange[1])
	isApartment := fmt.Sprintf("Is apartment : %t\n", fltr.isApartment)
	cty := fmt.Sprintf("Land area range : %s\n", fltr.city)
	delType := fmt.Sprintf("Deal type : %s\n", fltr.dealType)
	sourcSite := fmt.Sprintf("Source site : %s - %s\n", fltr.sourceSite[0], fltr.sourceSite[1])

	body := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s",
		report, priceRange, landAreaRange, roomRange, floorRange,
		haveStorage, publishDataRange, haveElevator,
		builtYearRange, isApartment, cty, delType, sourcSite)
	selector.Inline(
		selector.Row(room, landArea, price),
		selector.Row(publishData, storage, floor),
		selector.Row(apartment, builtYear, elevator),
		selector.Row(sourceSite, dealType, city),
		selector.Row(finish),
		selector.Row(back),
	)
	return c.Edit(
		body,
		selector,
	)
}

func roomBtn(c telebot.Context) error {
	roomRangeStatus = true
	message := "Enter your room range, separated by , :\nFor example \"1,2\""
	selector.Inline(
		selector.Row(roomRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func landAreaBtn(c telebot.Context) error {
	landAreaRangeStatus = true
	message := "Enter your land area range, separated by , :\nFor example \"100,200\""
	selector.Inline(
		selector.Row(landAreaRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func priceBtn(c telebot.Context) error {
	priceRangeStatus = true
	message := "Enter your price range, separated by , :\nFor example \"1000000,2000000\""
	selector.Inline(
		selector.Row(priceRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func publishDataBtn(c telebot.Context) error {
	publishDataRangeStatus = true
	message := "Enter your publish date range, separated by , :\nFor example \"2024-10-11,2024-12-01\""
	selector.Inline(
		selector.Row(publishDataRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func storageBtn(c telebot.Context) error {
	haveStorageStatus = true
	message := "Enter your storage status :\nFor example \"ok\""
	selector.Inline(
		selector.Row(storageRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func floorBtn(c telebot.Context) error {
	floorRangeStatus = true
	message := "Enter your floor range, separated by , :\nFor example \"2,3\""
	selector.Inline(
		selector.Row(floorRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func apartmentBtn(c telebot.Context) error {
	isApartmentStatus = true
	message := "Is it apartment :\nFor example \"No\""
	selector.Inline(
		selector.Row(apartmentRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func builtYearBtn(c telebot.Context) error {
	builtYearRangeStatus = true
	message := "Enter your build range, separated by , :\nFor example \"4,10\""
	selector.Inline(
		selector.Row(builtYearRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func elevatorBtn(c telebot.Context) error {
	haveElevatorStatus = true
	message := "Enter your elevetor status :\nFor example \"ok\""
	selector.Inline(
		selector.Row(elevatorRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func sourceSiteBtn(c telebot.Context) error {
	sourceSiteStatus = true
	message := "Enter your source type, separated by , :\nFor example \"divar,sheypoor\""
	selector.Inline(
		selector.Row(sourceSiteRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func dealTypeBtn(c telebot.Context) error {
	dealTypeStatus = true
	message := "Enter your dealing type, :\nFor example \"buy\""
	selector.Inline(
		selector.Row(dealTypeRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func cityBtn(c telebot.Context) error {
	cityStatus = true
	message := "Enter your city name  :\nFor example \"gazvin\""
	selector.Inline(
		selector.Row(cityRemove),
		selector.Row(backFilterEdit),
	)
	return c.Edit(
		message,
		selector,
	)
}
func roomRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.roomRange = []int{}
	roomRangeStatus = false
	return reportFilter(c)
}
func landAreaRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.landAreaRange = []int{}
	landAreaRangeStatus = false
	return reportFilter(c)
}
func priceRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.priceRange = []int{}
	priceRangeStatus = false
	return reportFilter(c)
}
func publishDataRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.publishDataRange = []string{}
	publishDataRangeStatus = false
	return reportFilter(c)
}
func storageRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.haveStorage = false
	haveStorageStatus = false
	return reportFilter(c)
}
func floorRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.floorRange = []int{}
	floorRangeStatus = false
	return reportFilter(c)

}
func apartmentRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.isApartment = false
	isApartmentStatus = false
	return reportFilter(c)
}
func builtYearRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.builtYearRange = []int{}
	builtYearRangeStatus = false
	return reportFilter(c)
}
func elevatorRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.haveElevator = false
	haveElevatorStatus = false
	return reportFilter(c)
}
func sourceSiteRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.sourceSite = []int{}
	sourceSiteStatus = false
	return reportFilter(c)
}
func dealTypeRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.dealType = rent
	dealTypeStatus = false
	return reportFilter(c)
}
func cityRemoveBtn(c telebot.Context) error {
	userId := c.Sender().ID
	fltr := filter[userId]
	fltr.city = ""
	cityStatus = false
	return reportFilter(c)

}
func finishBtn(c telebot.Context) error {
	loginMessage := fmt.Sprintf("Welcome dear %s %s", c.Chat().FirstName, c.Chat().LastName)
	selector.Inline(
		selector.Row(profile, bookmark, monitoring),
		selector.Row(post, export, report),
		selector.Row(logout),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func backFilterEditBtn(c telebot.Context) error {
	return reportFilter(c)
}

func selectReportBtn(c telebot.Context) error {
	message := fmt.Sprintf("report \"%s\" is selected , what operation do you want to performe ?", "vespa 200m")
	selector.Inline(
		selector.Row(deleteReport, editReport),
		selector.Row(report),
	)
	return c.Edit(
		message,
		selector,
	)
}

func deleteReportBtn(c telebot.Context) error {
	delete(filter, c.Sender().ID)
	return reportBtn(c)
}

func editReportBtn(c telebot.Context) error {
	delete(filter, c.Sender().ID)
	return newReportBtn(c)
}

func backPostBtn(c telebot.Context) error {
	return postBtn(c)
}

func backFilterBtn(c telebot.Context) error {
	loginMessage := fmt.Sprintf("Welcome dear %s %s", c.Chat().FirstName, c.Chat().LastName)
	selector.Inline(
		selector.Row(profile, bookmark, monitoring),
		selector.Row(post, export, report),
		selector.Row(logout),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func backBtn(c telebot.Context) error {
	selector.Inline(
		selector.Row(profile, bookmark, monitoring),
		selector.Row(post, export, report),
		selector.Row(logout),
	)

	return c.Edit(
		fmt.Sprintf("Welcome dear %s %s", c.Chat().FirstName, c.Chat().LastName),
		selector,
	)
}

func backMonitorAdminBtn(c telebot.Context) error {
	loginMessage := fmt.Sprintf("%s", "Error:")
	selector.Inline(
		selector.Row(monitoringAdmin, usr, accessToUserAccount),
		selector.Row(logout),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func backMonitorSuperAdminBtn(c telebot.Context) error {
	loginMessage := fmt.Sprintf("%s", "Error:")
	selector.Inline(
		selector.Row(monitoringSuperAdmin, setting, accessToUserAccount),
		selector.Row(admin, usr),
		selector.Row(logout),
	)
	return c.Edit(
		loginMessage,
		selector,
	)
}

func profileBtn(c telebot.Context) error {
	c.Delete()
	return c.Send("Hello mr mohammadi !")
}

func onText(c telebot.Context) error {
	if loginStatus {
		userID := c.Sender().ID
		c.Delete()
		if msg, ok := message[userID]; ok {
			_ = bot.Delete(msg)
		}
		if _, ok := data[userID]; ok && data[userID].userName == "" {
			data[userID] = user{
				userName: c.Text(),
				password: "",
			}

			resp, err := bot.Send(
				c.Sender(),
				"Please enter your password :",
			)
			if err != nil {
				return err
			}

			message[c.Sender().ID] = resp
		} else if _, ok := data[userID]; ok && data[userID].userName != "" && data[userID].password == "" {
			userName := data[userID].userName
			passWord := c.Text()
			data[userID] = user{
				userName: userName,
				password: passWord,
			}

			// is it user or admin or super admin :
			var loginMessage string

			switch data[userID].userType {
			case USER:
				loginMessage = fmt.Sprintf("Welcome dear %s %s", c.Chat().FirstName, c.Chat().LastName)
				selector.Inline(
					selector.Row(profile, bookmark, monitoring),
					selector.Row(post, export, report),
					selector.Row(logout),
				)
			case ADMIN:
				loginMessage = fmt.Sprintf("%s", "Error:")
				selector.Inline(
					selector.Row(monitoringAdmin, usr, accessToUserAccount),
					selector.Row(logout),
				)
			case SUPERADMIN:
				loginMessage = fmt.Sprintf("%s", "Error:")
				selector.Inline(
					selector.Row(monitoringSuperAdmin, setting, accessToUserAccount),
					selector.Row(admin, usr),
					selector.Row(logout),
				)
			}
			resp, err := bot.Send(
				c.Sender(),
				loginMessage,
				selector,
			)
			if err != nil {
				return err
			}

			message[userID] = resp
			loginStatus = false
		}
	} else if registerUserStatus {
		userID := c.Sender().ID
		c.Delete()
		if msg, ok := message[userID]; ok {
			_ = bot.Delete(msg)
		}
		if _, ok := data[userID]; ok && data[userID].userName == "" {
			data[userID] = user{
				userName: c.Text(),
				password: "",
			}

			resp, err := bot.Send(
				c.Sender(),
				"Please enter your password :",
			)
			if err != nil {
				return err
			}

			message[c.Sender().ID] = resp
		} else if _, ok := data[userID]; ok && data[userID].userName != "" && data[userID].password == "" {
			userName := data[userID].userName
			passWord := c.Text()
			data[userID] = user{
				userName: userName,
				password: passWord,
			}
			resp, err := bot.Send(
				c.Sender(),
				"Please enter your email to recive exports :",
			)
			if err != nil {
				return err
			}

			message[c.Sender().ID] = resp
		} else if _, ok := data[userID]; ok && data[userID].userName != "" && data[userID].password != "" && data[userID].email == "" {
			userName := data[userID].userName
			passWord := data[userID].password
			email := c.Text()
			data[userID] = user{
				userName: userName,
				password: passWord,
				email:    email,
				userType: USER,
			}
			selector.Inline(selector.Row(register, login))
			resp, err := bot.Send(
				c.Sender(),
				"Select an operation :",
				selector)
			if err != nil {
				return err
			}
			message[userID] = resp
			registerUserStatus = false
		}
	} else if registerAdminStatus {
		userID := c.Sender().ID
		c.Delete()
		if msg, ok := message[userID]; ok {
			_ = bot.Delete(msg)
		}
		if _, ok := data[userID]; ok && data[userID].userName == "" {
			data[userID] = user{
				userName: c.Text(),
				password: "",
			}

			resp, err := bot.Send(
				c.Sender(),
				"Please enter your password :",
			)
			if err != nil {
				return err
			}

			message[c.Sender().ID] = resp
		} else if _, ok := data[userID]; ok && data[userID].userName != "" && data[userID].password == "" {
			userName := data[userID].userName
			passWord := c.Text()
			data[userID] = user{
				userName: userName,
				password: passWord,
			}
			resp, err := bot.Send(
				c.Sender(),
				"Please enter your email to recive exports :",
			)
			if err != nil {
				return err
			}

			message[c.Sender().ID] = resp
		} else if _, ok := data[userID]; ok && data[userID].userName != "" && data[userID].password != "" && data[userID].email == "" {
			userName := data[userID].userName
			passWord := data[userID].password
			email := c.Text()
			data[userID] = user{
				userName: userName,
				password: passWord,
				email:    email,
				userType: ADMIN,
			}
			loginMessage := fmt.Sprintf("%s", "Error:")
			selector.Inline(
				selector.Row(monitoringSuperAdmin, setting, accessToUserAccount),
				selector.Row(admin, usr),
				selector.Row(logout),
			)
			resp, err := bot.Send(
				c.Sender(),
				loginMessage,
				selector,
			)
			if err != nil {
				return err
			}
			message[userID] = resp
			registerAdminStatus = false
		}
	} else if newReportStatus {
		userID := c.Sender().ID
		c.Delete()
		if msg, ok := message[userID]; ok {
			_ = bot.Delete(msg)
		}
		selector.Inline(
			selector.Row(newReportNo, newReportYes),
			selector.Row(newReportBack),
		)
		resp, err := bot.Send(
			c.Sender(),
			fmt.Sprintf("is it active ? \n\nTip : if it is active it starts working immediatly after creation . "),
			selector,
		)
		if err != nil {
			return err
		}
		message[userID] = resp
		newReportStatus = false
	} else if priceRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		prices := strings.Split(c.Text(), ",")
		fltr.priceRange[0], _ = strconv.Atoi(prices[0])
		fltr.priceRange[1], _ = strconv.Atoi(prices[1])
		priceRangeStatus = false
		return reportFilter(c)
	} else if landAreaRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		land := strings.Split(c.Text(), ",")
		fltr.landAreaRange[0], _ = strconv.Atoi(land[0])
		fltr.landAreaRange[1], _ = strconv.Atoi(land[1])
		landAreaRangeStatus = false
		return reportFilter(c)
	} else if roomRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		room := strings.Split(c.Text(), ",")
		fltr.roomRange[0], _ = strconv.Atoi(room[0])
		fltr.roomRange[1], _ = strconv.Atoi(room[1])
		roomRangeStatus = false
		return reportFilter(c)
	} else if roomRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		if c.Text() == "ok" {
			fltr.haveStorage = true
		} else {
			fltr.haveStorage = false
		}
		haveStorageStatus = false
		return reportFilter(c)
	} else if roomRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		floor := strings.Split(c.Text(), ",")
		fltr.floorRange[0], _ = strconv.Atoi(floor[0])
		fltr.floorRange[1], _ = strconv.Atoi(floor[1])
		floorRangeStatus = false
		return reportFilter(c)
	} else if isApartmentStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		if c.Text() == "ok" {
			fltr.isApartment = true
		} else {
			fltr.isApartment = false
		}
		isApartmentStatus = false
		return reportFilter(c)
	} else if builtYearRangeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		buildYear := strings.Split(c.Text(), ",")
		fltr.builtYearRange[0], _ = strconv.Atoi(buildYear[0])
		fltr.builtYearRange[1], _ = strconv.Atoi(buildYear[1])
		builtYearRangeStatus = false
		return reportFilter(c)
	} else if haveElevatorStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		if c.Text() == "ok" {
			fltr.haveElevator = true
		} else {
			fltr.haveElevator = false
		}
		haveElevatorStatus = false
		return reportFilter(c)
	} else if sourceSiteStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		sourceType := strings.Split(c.Text(), ",")
		fltr.sourceSite[0], _ = strconv.Atoi(sourceType[0])
		fltr.sourceSite[1], _ = strconv.Atoi(sourceType[1])
		sourceSiteStatus = false
		return reportFilter(c)
	} else if dealTypeStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		switch c.Text() {
		case "rent":
			fltr.dealType = rent
		case "buy":
			fltr.dealType = buy
		case "sell":
			fltr.dealType = sell
		}
		dealTypeStatus = false
		return reportFilter(c)
	} else if cityStatus {
		userID := c.Sender().ID
		c.Delete()
		fltr := filter[userID]
		fltr.city = c.Text()
		cityStatus = false
		return reportFilter(c)
	} else if newPostStatus {
		c.Delete()
		//..
		newPostStatus = false
		return postBtn(c)
	}
	return nil
}
