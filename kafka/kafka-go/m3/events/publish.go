package events

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	OrderReceivedTopic = "event"
)

func PublishEvent(OrderReceived OrderReceived, partition int) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", OrderReceivedTopic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(fmt.Sprintf("%v", OrderReceived))},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
		return
	}
}
