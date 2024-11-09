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

	profile    = selector.Data("Profile", "Profile")
	bookmark   = selector.Data("Bookmark", "Bookmark")
	monitoring = selector.Data("Monitoring", "Monitoring")
	post       = selector.Data("Post", "Post")
	export     = selector.Data("Export", "Export")
	report     = selector.Data("Report", "Report")
	logout     = selector.Data("Logout", "Logout")

	monitoringUser      = selector.Data("Monitoring", "MonitoringUser")
	usr                 = selector.Data("User", "User")
	accessToUserAccount = selector.Data("Access To User Account", "AccessToUserAccount")

	setting = selector.Data("Setting", "Setitng")
	admin   = selector.Data("Admin", "Admin")

	globalBookmark = selector.Data("Global Bookmarks", "Global_Bookmarks")
	myBookmark     = selector.Data("My Bookmarks", "My_Bookmarks")
	newPost        = selector.Data("Create New Post", "Create_new_post")
	back           = selector.Data("Back", "Back")
	refresh        = selector.Data("Refresh", "Refresh")
	back_repost    = selector.Data("Back", "Back_repost")
	skip           = selector.Data("Skip", "Skip")

	// reply btn
	cancle = menu.Text("Cancle")

	// state
	loginStatus    = false
	registerStatus = false
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
	_, err := bot.Send(c.Sender(), "Select an operation :", selector)
	return err
}

func registerBtn(c telebot.Context) error {
	c.Delete()
	registerStatus = true
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

func monitoringBtn(c telebot.Context) error {
	var (
		prevPageMonitor = selector.Data("Previuse page", "Previuse_page_monitoring")
		nexPageMonitor  = selector.Data("Next page", "Next_page_monitoring")
	)
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
	var (
		prevPageMyBkMark = selector.Data("Previuse page", "Previuse_page_my_bookmark")
		nexPageMyBkMark  = selector.Data("Next page", "Next_page_my_bookmark")
	)
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
	var (
		prevPageGlobBkMark = selector.Data("Previuse page", "Previuse_page_global_bookmark")
		nexPageGlobBkMark  = selector.Data("Next page", "Next_page_global_bookmark")
	)
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
					selector.Row(monitoringUser, usr, accessToUserAccount),
					selector.Row(logout),
				)
			case SUPERADMIN:
				loginMessage = fmt.Sprintf("%s", "Error:")
				selector.Inline(
					selector.Row(monitoringUser, setting, accessToUserAccount),
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
	} else if registerStatus {
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
			registerStatus = false
		}
	}
	return nil
}
