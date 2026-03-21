package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	c := NewContainer()

	asserts := assert.New(t)
	asserts.NotNil(c)
}

func TestContainer_GetHealthHandler(t *testing.T) {
	c := NewContainer()
	h := c.GetHealthHandler()

	asserts := assert.New(t)
	asserts.NotNil(h)
}

func TestContainer_GetSendMailHandler(t *testing.T) {
	c := NewContainer()
	h := c.GetSendMailHandler()

	asserts := assert.New(t)
	asserts.NotNil(h)
}
