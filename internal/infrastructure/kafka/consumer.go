package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func ListenToOrderEvents(brokers []string, topic string, notify func(event OrderEvent)) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  "telegram-bot-group", // общий для всех инстансов бота
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка чтения из Kafka: %v", err)
			continue
		}

		var event OrderEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Ошибка парсинга события: %v", err)
			continue
		}

		notify(event)
	}
}
