package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Producer(broker []string, topic string) {
	k := kafka.NewWriter(kafka.WriterConfig{
		Brokers: broker,
		Topic:   topic,
	})

	msg := kafka.Message{
		Key:   []byte("1"),
		Value: []byte("Hello Alexis"),
	}

	err := k.WriteMessages(context.Background(), msg)

	if err != nil {
		panic("could not write message " + err.Error())
	}
}

func Consumer(broker []string, topic string) {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers: broker,
		Topic:   topic,
	})

	for {
		msg, err := k.ReadMessage(context.Background())

		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
	}
}
