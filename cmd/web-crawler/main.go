package main

import (
	telegrambot "Golang-bc8-quera/web_crawler/pkg/telegram-bot"
	"fmt"
)

func main() {
	fmt.Println("Hello world !")
	go telegrambot.StartBot()

	var forever chan struct{}
	<-forever
}
