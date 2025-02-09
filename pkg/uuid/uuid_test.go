package uuid

import (
	"testing"

	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	uuid, err := New()
	logger.Info(uuid)
	assert.Nil(t, err)
	assert.NotEmpty(t, uuid)
	assert.Equal(t, 36, len(uuid))
	// unique
	uuid2, _ := New()
	assert.NotEqual(t, uuid, uuid2)
}
