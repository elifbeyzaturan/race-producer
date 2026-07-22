package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokerAddr string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerAddr),
		Topic:    "race-updates",
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{writer: writer}
}

func (p *Producer) Send(ctx context.Context, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = p.writer.WriteMessages(ctx, kafka.Message{Value: data})
	if err != nil {
		log.Printf("Failed to write message to Kafka: %v", err)
		return err
	}
	return nil
}

func (p *Producer) Close() {
	p.writer.Close()
}
