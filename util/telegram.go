package util

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/provalidator/stakescan-indexer/config"
)

func SendTelegramMsg(cfg config.Telegram, msg string) error {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return fmt.Errorf("new bot api: %w", err)
	}
	_, err = bot.Send(tgbotapi.NewMessage(cfg.ChatID, msg))
	if err != nil {
		return fmt.Errorf("bot.Send: %w", err)
	}
	return nil
}
