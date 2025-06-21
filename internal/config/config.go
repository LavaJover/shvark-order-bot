package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type OrderBotConfig struct {
	Env string 		   `yaml:"env"`
	GRPCServer 		   `yaml:"grpc_server"`
	OrderBotDB 		   `yaml:"order_db"`
	LogConfig 		   `yaml:"log_config"`
	BotToken	string `yanl:"bot_token"`
}

type GRPCServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type OrderBotDB struct {
	Dsn string `yaml:"dsn"`
}

type LogConfig struct {
	LogLevel 	string 	`yaml:"log_level"`
	LogFormat 	string 	`yaml:"log_format"`
	LogOutput 	string 	`yaml:"log_output"`
}

func MustLoad() *OrderBotConfig {

	// Processing env config variable and file
	configPath := os.Getenv("ORDER_BOT_CONFIG_PATH")

	if configPath == ""{
		log.Fatalf("ORDER_BOT_CONFIG_PATH was not found\n")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatalf("failed to find config file: %v\n", err)
	}

	// YAML to struct object
	var cfg OrderBotConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("failed to read config file: %v", err)
	}

	return &cfg
}