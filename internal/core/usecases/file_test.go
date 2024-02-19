package usecases

import (
	"context"
	"errors"
	"github.com/nunoonu/mail-consumer/internal/core/domain"
	"github.com/nunoonu/mail-consumer/internal/core/ports"
	"github.com/nunoonu/mail-consumer/internal/core/ports/mocks"
	"github.com/nunoonu/mail-consumer/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FileTestSuite struct {
	suite.Suite

	mailRepository     ports.MailRepository
	mockMailRepository *mocks.MailRepository
}

func (suite *FileTestSuite) SetupTest() {
	suite.mailRepository = repositories.NewMailRepository(nil)
	suite.mockMailRepository = &mocks.MailRepository{}
}

func TestUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}

func (suite *FileTestSuite) TestUploadSuite() {
	suite.Run("Success", func() {
		suite.mockMailRepository.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		usecase := NewFileUseCase(suite.mockMailRepository)
		actual, err := usecase.Send(context.TODO(), &domain.SendFileRequest{})
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), &domain.SendFileResponse{}, actual)
	})

	suite.Run("Error", func() {
		suite.mockMailRepository.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("couldn't be sent")).Once()

		usecase := NewFileUseCase(suite.mockMailRepository)
		_, err := usecase.Send(context.TODO(), &domain.SendFileRequest{})
		assert.Error(suite.T(), err)
	})

}
