package mail

import (
	"testing"

	"github.com/ijufumi/email-service/internal/infrastructure/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAmazonSESService(t *testing.T) {
	cfg := &config.Config{}
	svc := NewAmazonSESService(cfg)

	asserts := assert.New(t)
	asserts.NotNil(svc)
}
