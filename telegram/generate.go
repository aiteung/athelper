package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetPhone() tgbotapi.ReplyKeyboardMarkup {
	rowskey := make([]tgbotapi.KeyboardButton, 0, 1)
	keyboard := tgbotapi.KeyboardButton{
		Text:            "Get My Phone",
		RequestContact:  true,
		RequestLocation: true,
		RequestPoll:     nil,
	}
	rowskey = append(rowskey, keyboard)
	kyb := tgbotapi.NewReplyKeyboard(rowskey)
	kyb.OneTimeKeyboard = true
	fmt.Printf("Keyboard: %v", kyb)
	return kyb
}
