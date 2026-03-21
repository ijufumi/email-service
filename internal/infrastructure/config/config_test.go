package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad_DefaultValues(t *testing.T) {
	// Ensure relevant env vars are unset
	os.Unsetenv("MAIL_CHARSET")
	os.Unsetenv("MAIL_FROM_ADDRESS")

	c := Load()

	asserts := assert.New(t)
	asserts.NotNil(c)
	asserts.Equal("UTF-8", c.Mail.Charset)
	asserts.Equal("", c.Mail.FromAddress)
}

func TestLoad_WithEnvVars(t *testing.T) {
	t.Setenv("MAIL_CHARSET", "ISO-2022-JP")
	t.Setenv("MAIL_FROM_ADDRESS", "sender@example.com")
	t.Setenv("AWS_ACCESS_KEY_ID", "test-access-key")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test-secret-key")
	t.Setenv("AWS_REGION", "ap-northeast-1")
	t.Setenv("SENDGRID_API_KEY", "SG.test-api-key")

	c := Load()

	asserts := assert.New(t)
	asserts.NotNil(c)
	asserts.Equal("ISO-2022-JP", c.Mail.Charset)
	asserts.Equal("sender@example.com", c.Mail.FromAddress)
	asserts.Equal("test-access-key", c.Mail.SES.AwsAccessKeyID)
	asserts.Equal("test-secret-key", c.Mail.SES.AwsSecretAccessKEY)
	asserts.Equal("ap-northeast-1", c.Mail.SES.AwsRegion)
	asserts.Equal("SG.test-api-key", c.Mail.SendGrid.SendGridAPIKEY)
}
