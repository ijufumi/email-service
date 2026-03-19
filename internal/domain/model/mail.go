package model

// SendMailRequest represents a request to send an email
type SendMailRequest struct {
	ToAddress string
	Subject   string
	Body      string
}
