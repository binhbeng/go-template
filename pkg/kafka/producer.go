package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	Produce(ctx context.Context, topic string, message any) error
	Close() error
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string) KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &kafkaProducer{
		writer: writer,
	}
}

func (k *kafkaProducer) Produce(ctx context.Context, topic string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = k.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: body,
		Time:  time.Now(),
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (k *kafkaProducer) Close() error {
	if err := k.writer.Close(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}