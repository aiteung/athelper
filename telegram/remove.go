package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func RemoveKeyboard() (kyb tgbotapi.ReplyKeyboardRemove) {
	kyb.RemoveKeyboard = true
	return
}
