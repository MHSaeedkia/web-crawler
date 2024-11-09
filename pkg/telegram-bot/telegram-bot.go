package telegrambot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

type user struct {
	userName string
	password string
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

	globalBookmark = selector.Data("Global Bookmarks", "Global_Bookmarks")
	myBookmark     = selector.Data("My Bookmarks", "My_Bookmarks")
	back           = selector.Data("Back", "Back")
	refresh        = selector.Data("Refresh", "Refresh")

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

	bot.Handle(&back, backBtn)
	bot.Handle(&myBookmark, myBookmarkBtn)
	bot.Handle(&globalBookmark, globalBookmarkBtn)

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
		prevPageMonitor = selector.Data("Previuse page", "Previuse page monitoring")
		nexPageMonitor  = selector.Data("Next page", "Next page monitoring")
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
		prevPageMyBkMark = selector.Data("Previuse page", "Previuse page my bookmark")
		nexPageMyBkMark  = selector.Data("Next page", "Next page my bookmark")
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
		prevPageGlobBkMark = selector.Data("Previuse page", "Previuse page global bookmark")
		nexPageGlobBkMark  = selector.Data("Next page", "Next page global bookmark")
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

			selector.Inline(
				selector.Row(profile, bookmark, monitoring),
				selector.Row(post, export, report),
				selector.Row(logout),
			)
			resp, err := bot.Send(
				c.Sender(),
				fmt.Sprintf("Welcome dear %s %s", c.Chat().FirstName, c.Chat().LastName),
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
