package telegrambot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

var (
	menu     = &telebot.ReplyMarkup{ResizeKeyboard: true}
	selector = &telebot.ReplyMarkup{}

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

	// reply btn
	cancle = menu.Text("Cancle")
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

	bot, err := telebot.NewBot(pref)
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
	bot.Handle(&login, loginBtn)
	bot.Handle("/home", home)
	bot.Handle(&profile, profileBtn)

	// start bot
	bot.Start()
}

func start(c telebot.Context) error {
	// c.Delete()
	var (
		message = "Select an operation :"
	)
	selector.Inline(
		selector.Row(register, login),
	)
	return c.Send(message, selector)
}

func loginBtn(c telebot.Context) error {
	c.Delete()
	var (
		message = "Please enter your username :"
	)
	return c.Send(message)
}

func home(c telebot.Context) error {
	c.Delete()
	var (
		firstName = c.Chat().FirstName
		lastName  = c.Chat().LastName
		message   = fmt.Sprintf("Welcome dear %s %s", firstName, lastName)
	)

	selector.Inline(
		selector.Row(profile, bookmark, monitoring),
		selector.Row(post, export, report),
		selector.Row(logout),
	)
	return c.Send(message, selector)
}

func profileBtn(c telebot.Context) error {
	c.Delete()
	return c.Send("Hello mr mohammadi !")
}
