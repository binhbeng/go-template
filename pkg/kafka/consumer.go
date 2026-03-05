package kafka

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	Consume(ctx context.Context, handler func([]byte) error) error
	Close() error
}

type kafkaConsumer struct {
	reader *kafka.Reader
	logger *zerolog.Logger
}

func NewKafkaConsumer(brokers []string, topic string, groupID string, logger *zerolog.Logger) KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &kafkaConsumer{
		reader: reader,
		logger: logger,
	}
}

func (k *kafkaConsumer) Consume(ctx context.Context, handler func([]byte) error) error {
	go func() {
		for {
			msg, err := k.reader.FetchMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				k.logger.Error().Err(err).Msg("failed to fetch message")
				continue
			}

			if err := handler(msg.Value); err != nil {
				k.logger.Error().Err(err).Msg("handler failed")
				// không commit → sẽ retry
				continue
			}

			if err := k.reader.CommitMessages(ctx, msg); err != nil {
				k.logger.Error().Err(err).Msg("failed to commit message")
			}
		}
	}()

	return nil
}

func (k *kafkaConsumer) Close() error {
	if err := k.reader.Close(); err != nil {
		k.logger.Error().Err(err).Msg("failed to close consumer")
		return err
	}
	return nil
}