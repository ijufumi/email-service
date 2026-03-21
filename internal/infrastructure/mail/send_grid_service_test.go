package mail

import (
	"testing"

	"github.com/ijufumi/email-service/internal/infrastructure/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSendGridService(t *testing.T) {
	cfg := &config.Config{}
	svc := NewSendGridService(cfg)

	asserts := assert.New(t)
	asserts.NotNil(svc)
}
