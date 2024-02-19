package usecases

import (
	"context"
	"github.com/nunoonu/mail-consumer/internal/core/domain"
	"github.com/nunoonu/mail-consumer/internal/core/ports"
)

type fileUseCase struct {
	mailRepository ports.MailRepository
}

func NewFileUseCase(mailRepository ports.MailRepository) ports.FileUseCase {
	return &fileUseCase{
		mailRepository: mailRepository}
}

func (f fileUseCase) Send(ctx context.Context, req *domain.SendFileRequest) (*domain.SendFileResponse, error) {
	return &domain.SendFileResponse{}, f.mailRepository.Send(ctx, req.FileName, req.File)
}
