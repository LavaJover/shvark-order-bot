package telegram

import (
	"fmt"
	"strings"

	"github.com/LavaJover/shvark-order-bot/internal/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, authUC domain.AuthUsecase) {
	args := strings.Fields(msg.Text)
	if len(args) < 1 {
		return
	}

	switch args[0] {
	case "/start":
		if len(args) != 2 {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Используй: /start <токен>"))
			return
		}
		token := args[1]
		traderID, err := authUC.Authorize(msg.From.ID, token)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ошибка авторизации: " + err.Error()))
			return
		}
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf("Успешный вход. Trader %s", traderID)))
	default:
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Неизвестная команда"))
	}
}