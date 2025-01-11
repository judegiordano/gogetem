package logger

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(log.DebugLevel)
	lvl := GetLogLevel()
	assert.Equal(t, lvl, log.DebugLevel)
}
