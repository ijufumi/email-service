package mock

import (
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/domain/repository"
	"github.com/pkg/errors"
)

type mockAmazonSesService struct {
	hasError bool
}

// Send is mock of Send method
func (mock *mockAmazonSesService) Send(_ model.SendMailRequest) error {
	if mock.hasError {
		return errors.New("mock mockAmazonSesService error")
	}

	return nil
}

// NewMockAmazonSesService is factory method of Mock of AmazonSESService
func NewMockAmazonSesService(hasError bool) repository.SendMailVendor {
	return &mockAmazonSesService{hasError}
}
