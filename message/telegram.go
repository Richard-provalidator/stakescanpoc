package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stakescanpoc/log"
	"github.com/stakescanpoc/models"
)

func SendTelegramMsg(msgStr string) {
	botToken := models.Config.Telegram.BotToken
	chatID := int64(models.Config.Telegram.ChatID)

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Logger.Error.Panic(err)
	}

	// Create a message config
	msg := tgbotapi.NewMessage(chatID, msgStr)

	// Send the message
	_, err = bot.Send(msg)
	if err != nil {
		log.Logger.Error.Println(err)
	}
}
