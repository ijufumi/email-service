package usecase

import (
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/domain/repository"
	"go.uber.org/zap"
)

// SendMailService sends email
type SendMailService interface {
	Send(contents model.SendMailRequest) error
}

type sendMailService struct {
	vendors []repository.SendMailVendor
}

// Send allows sending email
func (service *sendMailService) Send(contents model.SendMailRequest) error {
	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	var err error
	for _, vendor := range service.vendors {
		err = vendor.Send(contents)
		if err == nil {
			break
		}
		logger.Warn("Failed to send message", zap.Error(err))
	}
	return err
}

// NewSendMailService is factory method of SendMailService
func NewSendMailService(vendors []repository.SendMailVendor) SendMailService {
	return &sendMailService{vendors: vendors}
}
