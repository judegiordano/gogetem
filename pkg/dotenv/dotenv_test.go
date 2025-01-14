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
	found := String("SHOULD_EXIST")
	assert.NotNil(t, *found)
	assert.Equal(t, *found, "value")
}

func TestGetNilString(t *testing.T) {
	os.Clearenv()
	found := String("DOES_NOT_EXIST")
	assert.Nil(t, found)
}

func TestGetInt(t *testing.T) {
	os.Clearenv()
	if err := os.Setenv("SHOULD_EXIST", "23"); err != nil {
		t.Errorf("error setting env: %v", err)
	}
	found := Int("SHOULD_EXIST")
	assert.NotNil(t, *found)
	assert.Equal(t, *found, 23)
}

func TestGetNilInt(t *testing.T) {
	os.Clearenv()
	found := Int("DOES_NOT_EXIST")
	assert.Nil(t, found)
}

func TestGetBool(t *testing.T) {
	os.Clearenv()
	if err := os.Setenv("SHOULD_EXIST", "true"); err != nil {
		t.Errorf("error setting env: %v", err)
	}
	found := Bool("SHOULD_EXIST")
	assert.NotNil(t, *found)
	assert.Equal(t, *found, true)
}

func TestGetNilBool(t *testing.T) {
	os.Clearenv()
	found := Bool("DOES_NOT_EXIST")
	assert.Nil(t, found)
}
