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

	prevPage = selector.Data("Previuse page", "Previuse page")
	nexPage  = selector.Data("Next page", "Next page")
	back     = selector.Data("Back", "Back")
	refresh  = selector.Data("Refresh", "Refresh")

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
	bot.Handle(&back, backBtn)

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
	c.Delete()
	userId := c.Sender().ID
	page := selector.Data(fmt.Sprintf("page %v of %v", 1, 1), "page")
	log := fmt.Sprintf("crawler logs activities : \n%s", "log ..")
	selector.Inline(
		selector.Row(prevPage, page, nexPage),
		selector.Row(refresh),
		selector.Row(back),
	)
	resp, err := bot.Send(
		c.Sender(),
		log,
		selector,
	)
	if err != nil {
		return err
	}
	message[userId] = resp
	return nil
}

func backBtn(c telebot.Context) error {
	c.Delete()
	userId := c.Sender().ID
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

	message[userId] = resp
	return nil
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
