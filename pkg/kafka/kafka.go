package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaTopic = "responses-log"
	kafkaURL   = "localhost:9092"
)

func GetKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func GetKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaURL},
		Topic:     kafkaTopic,
		Partition: 0,
	})
}

func AppendCommandLog(ctx context.Context, kafkaWriter *kafka.Writer, k []byte, v []byte) error {
	msg := kafka.Message{
		// Topic: kafkaTopic,
		Key:   k,
		Value: v,
	}
	err := kafkaWriter.WriteMessages(ctx, msg)
	return err
}
