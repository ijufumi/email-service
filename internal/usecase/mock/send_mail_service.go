package mock

import (
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/usecase"
	"github.com/pkg/errors"
)

type mockSendMailService struct {
	hasError bool
}

// Send is mock of Send method
func (mock *mockSendMailService) Send(_ model.SendMailRequest) error {
	if mock.hasError {
		return errors.New("mock service error")
	}
	return nil
}

// NewMockSendMailService is factory method of Mock of SendMailService
func NewMockSendMailService(withError bool) usecase.SendMailService {
	return &mockSendMailService{hasError: withError}
}
