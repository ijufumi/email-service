package container

import (
	"github.com/ijufumi/email-service/internal/adapter/handler"
	"github.com/ijufumi/email-service/internal/domain/repository"
	infraconfig "github.com/ijufumi/email-service/internal/infrastructure/config"
	inframail "github.com/ijufumi/email-service/internal/infrastructure/mail"
	"github.com/ijufumi/email-service/internal/usecase"
	"go.uber.org/dig"
)

// Container is DI container
type Container interface {
	GetHealthHandler() handler.HealthHandler
	GetSendMailHandler() handler.SendMailHandler
}

type container struct {
	container *dig.Container
}

func (c *container) provide() {
	// config
	_ = c.container.Provide(infraconfig.Load)

	// mail vendors
	_ = c.container.Provide(func(cfg *infraconfig.Config) []repository.SendMailVendor {
		return []repository.SendMailVendor{
			inframail.NewAmazonSESService(cfg),
			inframail.NewSendGridService(cfg),
		}
	})

	// usecase
	_ = c.container.Provide(usecase.NewSendMailService)

	// handler
	_ = c.container.Provide(handler.NewHealthHandler)
	_ = c.container.Provide(handler.NewSendMailHandler)
}

// GetHealthHandler returns instance of HealthHandler
func (c *container) GetHealthHandler() handler.HealthHandler {
	var healthHandler handler.HealthHandler
	_ = c.container.Invoke(func(_healthHandler handler.HealthHandler) {
		healthHandler = _healthHandler
	})
	return healthHandler
}

// GetSendMailHandler returns instance of SendMailHandler
func (c *container) GetSendMailHandler() handler.SendMailHandler {
	var sendMailHandler handler.SendMailHandler
	_ = c.container.Invoke(func(_sendMailHandler handler.SendMailHandler) {
		sendMailHandler = _sendMailHandler
	})
	return sendMailHandler
}

// NewContainer creates instance of Container
func NewContainer() Container {
	c := container{container: dig.New()}
	c.provide()
	return &c
}
