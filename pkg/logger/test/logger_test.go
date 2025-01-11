package test

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/judegiordano/gogetem/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	logger.SetLogLevel(log.DebugLevel)
	lvl := logger.GetLogLevel()
	assert.Equal(t, lvl, log.DebugLevel)
}
