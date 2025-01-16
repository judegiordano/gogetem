package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	SetLogLevel(logrus.DebugLevel)
	lvl := GetLogLevel()
	Debug("you should see this :)")
	assert.Equal(t, lvl, logrus.DebugLevel)
}
