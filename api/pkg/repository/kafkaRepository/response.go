package kafkaRepository

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func KafkaResponse(write *kafka.Writer, errFunc, op string) {
	message := kafka.Message{
		Key:   []byte(op),
		Value: []byte(errFunc),
	}

	err := write.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Println("error writing to kafka:", err)
	}
	return
}
