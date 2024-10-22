package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// Создаем новый reader (консумер) для чтения сообщений из Kafka
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "TelegramShop",
		GroupID: "example-group",
	})
	defer reader.Close()

	fmt.Println("consumer running")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Ошибка при чтении сообщения:", err)
		}
		log.Printf("Получено сообщение: Функция: %s,  %s\n", string(msg.Key), string(msg.Value))
	}
}

//app:
//build: .
//
//ports:
//- "8000:8000"
//command: go run ./api/cmd/main.go
