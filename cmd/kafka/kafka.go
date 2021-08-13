package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Producer(broker []string, topic string, body string) {
	k := kafka.NewWriter(kafka.WriterConfig{
		Brokers: broker,
		Topic:   topic,
	})

	msg := kafka.Message{
		Key:   []byte("1"),
		Value: []byte(body),
	}

	err := k.WriteMessages(context.Background(), msg)

	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Consumer(broker []string, topic string) string {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     broker,
		Topic:       topic,
		GroupID:     "1",
		StartOffset: kafka.LastOffset,
	})

	for {
		msg, err := k.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Print(string(msg.Value))
	}
}
