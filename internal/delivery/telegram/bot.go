package telegram

import (
	"log"

	"github.com/LavaJover/shvark-order-bot/internal/domain"
	"github.com/LavaJover/shvark-order-bot/internal/infrastructure/kafka"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


type Bot struct {
	api *tgbotapi.BotAPI
	authUC domain.AuthUsecase
	orderChan chan domain.OrderNotification
}

func NewBot(botToken string, authUC domain.AuthUsecase) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &Bot{
		api: api,
		authUC: authUC,
		orderChan: make(chan domain.OrderNotification, 100),
	}, nil
}

func (b *Bot) Start(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	go b.listenForNotifications()

	for update := range updates {
		if update.Message != nil {
			handleMessage(b.api, update.Message, b.authUC)
		}
	}
}

func (b *Bot) Notify(event kafka.OrderEvent) {
	order := domain.OrderNotification{
		OrderID: event.OrderID,
		TraderID: event.TraderID,
		Status: event.Status,
		Amount: event.AmountFiat,
		Currency: event.Currency,
		BankName: event.BankName,
		Phone: event.Phone,
		CardNumber: event.CardNumber,
		Owner: event.Owner,
	}
	b.orderChan <- order
}

func (b *Bot) listenForNotifications() {
	for order := range b.orderChan {
		telegramIDs, err := b.authUC.GetTelegramIDsByTraderID(order.TraderID)
		if err != nil {
			log.Printf("Failed to get telegram bindings for trader %s\n", order.TraderID)
		}
		for _, telegramID := range telegramIDs {
			text := order.String()
			msg := tgbotapi.NewMessage(telegramID, text)
			_, err := b.api.Send(msg)
			if err != nil {
				log.Printf("Failed to send TG message: %v\n", err)
			}
		}
	}
}