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
	if err := os.Setenv("SHOULD_EXIST", "value"); err != nil {
		t.Errorf("error setting env: %v", err)
	}
	Load()
	found := String("SHOULD_EXIST")
	assert.NotNil(t, *found)
	assert.Equal(t, *found, "value")
}

func TestGetNilString(t *testing.T) {
	os.Clearenv()
	Load()
	found := String("DOES_NOT_EXIST")
	assert.Nil(t, found)
}
