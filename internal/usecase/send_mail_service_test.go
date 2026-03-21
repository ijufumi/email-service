package usecase

import (
  "github.com/ijufumi/email-service/internal/domain/model"
  "github.com/ijufumi/email-service/internal/domain/repository"
  "github.com/ijufumi/email-service/internal/infrastructure/mail/mock"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSendMailService_Send(t *testing.T) {
  mockAmazonSesService := mock.NewMockAmazonSesService(false)
  mockSendGridService := mock.NewMockSendGridService(false)

  sendMailService := NewSendMailService([]repository.SendMailVendor{mockAmazonSesService, mockSendGridService})
  err := sendMailService.Send(model.SendMailRequest{})

  asserts := assert.New(t)

  asserts.Nil(err)
}

func TestSendMailService_Send_Second_Success(t *testing.T) {
  mockAmazonSesService := mock.NewMockAmazonSesService(true)
  mockSendGridService := mock.NewMockSendGridService(false)

  sendMailService := NewSendMailService([]repository.SendMailVendor{mockAmazonSesService, mockSendGridService})
  err := sendMailService.Send(model.SendMailRequest{})

  asserts := assert.New(t)

  asserts.Nil(err)
}

func TestSendMailService_First_Success(t *testing.T) {
  mockAmazonSesService := mock.NewMockAmazonSesService(false)
  mockSendGridService := mock.NewMockSendGridService(true)

  sendMailService := NewSendMailService([]repository.SendMailVendor{mockAmazonSesService, mockSendGridService})
  err := sendMailService.Send(model.SendMailRequest{})

  asserts := assert.New(t)

  asserts.Nil(err)
}

func TestSendMailService_Error(t *testing.T) {
  mockAmazonSesService := mock.NewMockAmazonSesService(true)
  mockSendGridService := mock.NewMockSendGridService(true)

  sendMailService := NewSendMailService([]repository.SendMailVendor{mockAmazonSesService, mockSendGridService})
  err := sendMailService.Send(model.SendMailRequest{})

  asserts := assert.New(t)

  asserts.NotNil(err)
  asserts.Equal(err.Error(), "mock mockSendGridService error")
}

func TestSendMailService_Send_EmptyVendors(t *testing.T) {
  sendMailService := NewSendMailService([]repository.SendMailVendor{})
  err := sendMailService.Send(model.SendMailRequest{})

  asserts := assert.New(t)

  asserts.Nil(err)
}
