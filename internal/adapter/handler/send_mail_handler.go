package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ijufumi/email-service/internal/adapter/handler/dto"
	"github.com/ijufumi/email-service/internal/domain/model"
	"github.com/ijufumi/email-service/internal/usecase"
	"net/http"
)

// SendMailHandler is the handler for sending mail
type SendMailHandler interface {
	Send(ctx *gin.Context)
}

type sendMailHandler struct {
	sendMailService usecase.SendMailService
}

// Send allows sending email via Email Vendors
func (handler *sendMailHandler) Send(ctx *gin.Context) {
	var req dto.SendMailRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Result{Status: false, Error: &err})
		return
	}
	contents := model.SendMailRequest{
		ToAddress: req.ToAddress,
		Subject:   req.Subject,
		Body:      req.Body,
	}
	if err := handler.sendMailService.Send(contents); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Result{Status: false, Error: &err})
		return
	}
	ctx.JSON(http.StatusOK, dto.Result{Status: true})
}

// NewSendMailHandler is factory method of SendMailHandler
func NewSendMailHandler(sendMailService usecase.SendMailService) SendMailHandler {
	return &sendMailHandler{sendMailService: sendMailService}
}
