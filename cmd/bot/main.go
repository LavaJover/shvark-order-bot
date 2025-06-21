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
	fmt.Println(cfg)

	ssoAddr := "localhost:50051"
	ssoClient, err := grpcapi.NewSSOClient(ssoAddr)
	if err != nil {
		log.Fatalf("failed to connect SSO-client")
	}

	db := postgres.InitDB(cfg)
	authRepo := postgres.NewDefaultAuthRepository(db)
	authUC := usecase.NewAuthUsecase(authRepo, ssoClient)

	bot, err := telegram.NewBot("7096257833:AAEDRlZGucm_5-g0MxK4BiUqZ1bJo_Bon3M", authUC)
	if err != nil {
		log.Fatalf("failed to init bot")
	}

	go kafka.ListenToOrderEvents([]string{"localhost:9092"}, "order-events", bot.Notify)

	bot.Start()

}