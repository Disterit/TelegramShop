package kafkaHandler

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func KafkaResponse(writer *kafka.Writer, errFunc, op string) {
	msg := kafka.Message{
		Key:   []byte(op),
		Value: []byte(errFunc),
	}

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		fmt.Println("error writing to kafka:", err)
	}
	return
}
