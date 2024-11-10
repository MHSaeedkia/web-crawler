package telegrambot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

const (
	USER       = "User"
	ADMIN      = "Admin"
	SUPERADMIN = "SuperAdmin"
)

type user struct {
	userName string
	password string
	userType string
	email    string
}

var (
	bot *telebot.Bot

	menu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	selector = &telebot.ReplyMarkup{RemoveKeyboard: true}

	data    = make(map[int64]user)
	message = make(map[int64]*telebot.Message)

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
	post            = selector.Data("Post", "Post")
	export          = selector.Data("Export", "Export")
	report          = selector.Data("Report", "Report")
	logout          = selector.Data("Logout", "Logout")

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
	newPost        = selector.Data("Create New Post", "Create_New_Post")
	back           = selector.Data("Back", "Back")
	refresh        = selector.Data("Refresh", "Refresh")
	back_repost    = selector.Data("Back", "Back_Repost")
	skip           = selector.Data("Skip", "Skip")

	// reply btn
	cancle = menu.Text("Cancle")

	// state
	loginStatus         = false
	registerUserStatus  = false
	registerAdminStatus = false
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
	bot.Handle(&report, reportBtn)

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
	var (
		prevPagePost = selector.Data("Previuse page", "Previuse_page_post")
		nexPagePost  = selector.Data("Next page", "Next_page_post")
	)
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

func newPostBtn(c telebot.Context) error {
	posts := "Enter a name for your post : "

	// copule of inline btn .

	selector.Inline(
		selector.Row(newPost),
		selector.Row(back),
	)

	return c.Edit(
		posts,
		selector,
	)
}

func reportBtn(c telebot.Context) error {
	var (
		prevPageReport = selector.Data("Previuse page", "Previuse_page_report")
		nexPageReport  = selector.Data("Next page", "Next_page_report")
	)
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	report := fmt.Sprintf("The report \"%s\" has been selected , select a post to view , edit or delere :", "vespa 200 divar and sheypoor")

	// copule of inline btn .

	selector.Inline(
		selector.Row(prevPageReport, page, nexPageReport),
		selector.Row(newPost),
		selector.Row(back),
	)

	return c.Edit(
		report,
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
				selector)
			if err != nil {
				return err
			}
			message[userID] = resp
			registerAdminStatus = false
		}
	}
	return nil
}
