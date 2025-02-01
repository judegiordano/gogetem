package gravatar

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	email := "mail@mail.com"
	hash := GenerateHash(email)
	assert.Equal(t, hash, "7905d373cfab2e0fda04b9e7acc8c879")
	url := Url(email, "2048")
	assert.Equal(t, url, fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro&s=2048", hash))
}

func TestRequest(t *testing.T) {
	email := "mail@mail.com"
	bytes, err := Image(email, 2048)

	// should build bytes
	assert.Nil(t, err)
	assert.True(t, len(bytes) >= 1)

	path := os.TempDir()
	path = fmt.Sprint(filepath.Join(path, "avatar.png"))

	// should be able to save as img
	err = os.WriteFile(path, bytes, 0644)
	assert.Nil(t, err)

	// img should exist
	img, err := os.ReadFile(path)
	assert.Nil(t, err)
	assert.True(t, len(img) >= 1)

	// should remove path
	err = os.Remove(path)
	assert.Nil(t, err)

	// path should not exist
	missing, err := os.ReadFile(path)
	assert.NotNil(t, err)
	assert.Nil(t, missing)
}

func TestGenerate(t *testing.T) {
	gravatar := Gravatar{
		Email:   "mail@mail.com",
		Size:    2048,
		Default: Identicon,
	}
	data, err := gravatar.Generate()
	assert.NotNil(t, data)
	assert.Nil(t, err)

	path := fmt.Sprint(filepath.Join("./", "avatar.png"))

	// should be able to save as img
	err = os.WriteFile(path, data.Image, 0644)
	assert.Nil(t, err)

	// img should exist
	img, err := os.ReadFile(path)
	assert.Nil(t, err)
	assert.True(t, len(img) >= 1)

	// should remove path
	err = os.Remove(path)
	assert.Nil(t, err)

	// path should not exist
	missing, err := os.ReadFile(path)
	assert.NotNil(t, err)
	assert.Nil(t, missing)
}
