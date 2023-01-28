package main

import (
	"context"
	"fmt"

	"github.com/nikoksr/notify"
	telegram_alert_bot "github.com/nikoksr/notify/service/telegram"
)

var bot_token string
var bot_chat_ID int64

func init() {
	fmt.Println("Init value")

	bot_token = "5698880402:AAHDIojvC9vuuW9Ti0BD2zw-ASSmQmoI-Uo"
	bot_chat_ID = 5018541524
}

func main() {

	fmt.Println("Send Test Message")
	// Create a telegram service. Ignoring error for demo simplicity.
	telegramService, err := telegram_alert_bot.New(bot_token)
	if err != nil {
		fmt.Println(err)
	}

	// Passing a telegram chat id as receiver for our messages.
	// Basically where should our message be sent?
	telegramService.AddReceivers(bot_chat_ID)

	// Tell our notifier to use the telegram service. You can repeat the above process
	// for as many services as you like and just tell the notifier to use them.
	// Inspired by http middleware used in higher level libraries.
	notify.UseServices(telegramService)

	// Send a test message.
	send_err := notify.Send(
		context.Background(),
		"Subject/Title",
		"The actual message - Hello, you awesome gophers! :)",
	)

	if send_err != nil {
		fmt.Println(send_err)
	}
}
