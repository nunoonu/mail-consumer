package handlers

import (
	"context"
	"encoding/json"
	"github.com/nunoonu/mail-consumer/internal/core/domain"
	"github.com/nunoonu/mail-consumer/internal/core/ports"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type Mail struct {
	FileName string
	File     []byte
}

type FileHandler struct {
	auc ports.FileUseCase
}

func NewFileHandler(auc ports.FileUseCase) *FileHandler {
	return &FileHandler{auc: auc}
}

func (h *FileHandler) Do(ctx context.Context, mes kafka.Message) error {
	m := Mail{}
	err := json.Unmarshal(mes.Value, &m)
	if err != nil {
		slog.Error("failed to unmarshal message", slog.String("Err", err.Error()))
		return err
	}
	_, err = h.auc.Send(ctx, &domain.SendFileRequest{File: m.File, FileName: m.FileName})
	return err
}
