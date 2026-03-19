package repository

import "github.com/ijufumi/email-service/internal/domain/model"

// SendMailVendor is interface of mail vendor
type SendMailVendor interface {
	Send(contents model.SendMailRequest) error
}
