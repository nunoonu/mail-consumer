package ports

import (
	"context"
	"github.com/nunoonu/mail-consumer/internal/core/domain"
)

type FileUseCase interface {
	Send(ctx context.Context, req *domain.SendFileRequest) (*domain.SendFileResponse, error)
}
