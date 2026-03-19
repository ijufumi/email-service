package dto

// SendMailRequest is request parameter for send mail
type SendMailRequest struct {
	ToAddress string `json:"to_address" binding:"required,email"`
	Subject   string `json:"subject" binding:"required"`
	Body      string `json:"body" binding:"required"`
}
