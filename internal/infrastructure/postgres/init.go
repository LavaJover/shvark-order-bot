package postgres

import (
	"log"

	"github.com/LavaJover/shvark-order-bot/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.OrderBotConfig) *gorm.DB{
	dsn := cfg.OrderBotDB.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init db")
	}

	db.AutoMigrate(&TelegramBinding{})

	return db
}