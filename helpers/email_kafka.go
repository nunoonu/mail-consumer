package helpers

import (
	"github.com/segmentio/kafka-go"
)

type MailParams struct {
	URL   string
	Topic string
	Group string
}

func NewMailKafkaParams() *MailParams {
	return &MailParams{
		URL:   "localhost:9092",
		Topic: "email",
		Group: "email-group",
	}
}

func NewKafka(param *MailParams) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{param.URL},
		GroupID:  param.Group,
		Topic:    param.Topic,
		MaxBytes: 10e6, // 10MB
	})
}
