package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func SendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	log.Printf("\nSending Message %+v\n", msg)
	_, er := bot.Send(msg)
	if er != nil {
		log.Printf("\nError terjadi : %s\n", er.Error())
	}
}
func CreateMessage(userId int64, msgId int, msg string) (res tgbotapi.MessageConfig) {
	res = tgbotapi.NewMessage(userId, msg)
	res.ReplyMarkup = RemoveKeyboard()
	res.ReplyToMessageID = msgId
	fmt.Printf("\nCreated Template : %+v\n", res)
	return
}
