package mock

import (
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/domain/repository"
	"github.com/pkg/errors"
)

type mockSendGridService struct {
	hasError bool
}

// Send is mock of Send method
func (mock *mockSendGridService) Send(_ model.SendMailRequest) error {
	if mock.hasError {
		return errors.New("mock mockSendGridService error")
	}

	return nil
}

// NewMockSendGridService is factory method of Mock of SendGridService
func NewMockSendGridService(hasError bool) repository.SendMailVendor {
	return &mockSendGridService{hasError}
}
