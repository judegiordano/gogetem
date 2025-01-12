package dotenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeString(t *testing.T) {
	os.Clearenv()
	normalized := normalizeKey(" kEy ")
	assert.Equal(t, normalized, "KEY")
}

func TestGetString(t *testing.T) {
	os.Clearenv()
	if err := os.Setenv("should_exist", "value"); err != nil {
		t.Errorf("error setting env: %v", err)
	}
	Load()
	found := String("should_exist")
	assert.NotNil(t, found)
	assert.Equal(t, *found, "value")
}

func TestGetNilString(t *testing.T) {
	os.Clearenv()
	Load()
	found := String("does_not_exist")
	assert.Nil(t, found)
}
