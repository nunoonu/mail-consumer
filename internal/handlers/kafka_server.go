package handlers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log/slog"
	"strconv"
)

func NewKafkaService(kr *kafka.Reader, hdl *FileHandler) *KafkaService {
	return &KafkaService{
		Reader:      kr,
		FileHandler: hdl,
	}
}

type KafkaService struct {
	*kafka.Reader
	*FileHandler
}

func (h *KafkaService) Start() {
	for {
		ctx := context.Background()
		m, err := h.ReadMessage(context.Background())
		if err != nil {
			slog.Error("failed to read message", slog.String("Err", err.Error()))
			break
		}
		slog.Info("Received message", slog.Group("message",
			slog.String("topic", m.Topic),
			slog.String("partition", strconv.Itoa(m.Partition)),
			slog.String("offset", strconv.FormatInt(m.Offset, 10)),
		))
		if err = h.FileHandler.Do(ctx, m); err != nil {
			slog.Error("failed to process message", slog.String("Err", err.Error()))
		}
	}
}

func (h *KafkaService) Close() error {
	return h.Reader.Close()
}
