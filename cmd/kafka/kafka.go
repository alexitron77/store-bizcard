package kafka

import (
	"context"
	"fmt"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

var i = 0

func Producer(broker []string, topic string, body string) {
	k := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  broker,
		Topic:    topic,
		Balancer: &kafka.Hash{},
	})

	msg := kafka.Message{
		Key:   []byte(strconv.Itoa(i)),
		Value: []byte(body),
	}

	err := k.WriteMessages(context.Background(), msg)

	if err != nil {
		panic("could not write message " + err.Error())
	}

	// Increment i to write message to random partition
	i++
}

func Consumer(broker []string, topic string) string {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     broker,
		Topic:       topic,
		GroupID:     "5",
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
