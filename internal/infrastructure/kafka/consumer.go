package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

type KafkaConfig struct {
	Brokers    []string
    Topic      string
    Username   string
    Password   string
    Mechanism  string // "PLAIN", "SCRAM-SHA-256", etc.
    TLSEnabled bool
}

func NewKafkaConsumer(cfg KafkaConfig) (*KafkaConsumer, error) {
	readerConfig := kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic: cfg.Topic,
		GroupID: "telegram-bot-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	}

	if cfg.Username != "" && cfg.Password != "" {
        mechanism, err := createSASLMechanism(cfg.Mechanism, cfg.Username, cfg.Password)
        if err != nil {
            return nil, err
        }

        readerConfig.Dialer = &kafka.Dialer{
            SASLMechanism: mechanism,
            TLS:           createTLSConfig(cfg.TLSEnabled),
        }
    }

	return &KafkaConsumer{
        reader: kafka.NewReader(readerConfig),
    }, nil
}

func createSASLMechanism(mechanism, username, password string) (sasl.Mechanism, error) {
    switch mechanism {
    case "SCRAM-SHA-256":
        return scram.Mechanism(scram.SHA256, username, password)
    case "SCRAM-SHA-512":
        return scram.Mechanism(scram.SHA512, username, password)
    case "PLAIN":
        return plain.Mechanism{
            Username: username,
            Password: password,
        }, nil
    default:
        return nil, fmt.Errorf("unsupported SASL mechanism: %s", mechanism)
    }
}

func createTLSConfig(enabled bool) *tls.Config {
    if !enabled {
        return nil
    }
    return &tls.Config{
        InsecureSkipVerify: false, // В продакшене должно быть false
    }
}

func (consumer *KafkaConsumer) ListenToOrderEvents(notify func(event OrderEvent)) {
	slog.Info("start listening kafka messages")
	for {
		msg, err := consumer.reader.ReadMessage(context.Background())
		if err != nil {
			slog.Error("failed to read from kafka", "error", err.Error())
			continue
		}
		slog.Info("received kafka message")

		var event OrderEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("failed to parse order event", "error", err.Error())
			continue
		}

		notify(event)
	}
}
