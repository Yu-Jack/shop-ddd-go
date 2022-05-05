package dddcore

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

var (
	topic        = os.Getenv("KAFKA_TOPIC")
	partition, _ = strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	server       = os.Getenv("KAFKA_ENDPOINT")
)

type EventBus struct {
	conn   *kafka.Conn
	writer *kafka.Writer
}

func NewEventBus() *EventBus {
	conn, err := kafka.DialLeader(context.Background(), "tcp", server, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(server),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &EventBus{
		conn:   conn,
		writer: writer,
	}
}

func (eb *EventBus) getReader(groupID string) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{server},
		Topic:     topic,
		Partition: partition,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		GroupID:   groupID,
	})
	return reader
}

func (eb *EventBus) Subscribe(groupId string, cb func(value string)) {
	reader := eb.getReader(groupId)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		if groupId == string(m.Key) {
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			cb(string(m.Value))
		}
	}
}

func (eb *EventBus) Publish(evs []Event) {
	for _, ev := range evs {
		eb.writer.WriteMessages(
			context.Background(),
			kafka.Message{Key: []byte(ev.EventName), Value: ev.RawData},
		)
	}
}
