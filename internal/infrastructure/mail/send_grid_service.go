package mail

import (
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/domain/repository"
	"github.com/ijufumi/email-service/internal/infrastructure/config"
	"github.com/sendgrid/sendgrid-go"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridService struct {
	config *config.Config
}

// Send allows sending mail via SendGrid
func (service *sendGridService) Send(contents model.SendMailRequest) error {
	toAddress := sgmail.NewEmail(contents.ToAddress, contents.ToAddress)
	fromAddress := sgmail.NewEmail(service.config.Mail.FromAddress, service.config.Mail.FromAddress)
	message := sgmail.NewSingleEmailPlainText(fromAddress, contents.Subject, toAddress, contents.Body)
	client := service.createClient()
	_, err := client.Send(message)
	return err
}

func (service *sendGridService) createClient() *sendgrid.Client {
	return sendgrid.NewSendClient(service.config.Mail.SendGrid.SendGridAPIKEY)
}

// NewSendGridService is factory method of SendGridService
func NewSendGridService(config *config.Config) repository.SendMailVendor {
	return &sendGridService{
		config: config,
	}
}
