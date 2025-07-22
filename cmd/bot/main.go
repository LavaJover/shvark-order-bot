package main

import (
	"fmt"
	"log"

	"github.com/LavaJover/shvark-order-bot/internal/config"
	"github.com/LavaJover/shvark-order-bot/internal/delivery/telegram"
	"github.com/LavaJover/shvark-order-bot/internal/grpcapi"
	"github.com/LavaJover/shvark-order-bot/internal/infrastructure/kafka"
	"github.com/LavaJover/shvark-order-bot/internal/infrastructure/postgres"
	"github.com/LavaJover/shvark-order-bot/internal/usecase"
	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env")
	}
	// read config
	cfg := config.MustLoad()

	ssoAddr := fmt.Sprintf("%s:%s", cfg.SSOService.Host, cfg.SSOService.Port)
	ssoClient, err := grpcapi.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to connect SSO-client")
	}

	db := postgres.InitDB(cfg)
	authRepo := postgres.NewDefaultAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo, ssoClient)

	bot, err := telegram.NewBot(cfg.BotToken, authUC)
	if err != nil {
		log.Fatalf("failed to init bot")
	}

	go kafka.ListenToOrderEvents([]string{fmt.Sprintf("%s:%s", cfg.KafkaService.Host, cfg.KafkaService.Port)}, "order-events", bot.Notify)

	bot.Start()

}