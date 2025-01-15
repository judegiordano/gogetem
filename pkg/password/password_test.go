package password

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	original := "password"
	pwd, err := Hash(original)
	assert.Nil(t, err)
	assert.NotEqual(t, pwd, original)
}

func TestVaild(t *testing.T) {
	original := "password"
	hash, err := Hash(original)
	assert.Nil(t, err)
	assert.NotEqual(t, hash, original)

	match, err := Verify(original, hash)
	assert.Nil(t, err)
	assert.True(t, match)
}

func TestNotValid(t *testing.T) {
	original := "Password"
	hash, err := Hash(original)
	assert.Nil(t, err)
	assert.NotEqual(t, hash, original)

	match, err := Verify("password", hash)
	assert.Nil(t, err)
	assert.False(t, match)
}

func TestInvalidVersion(t *testing.T) {
	hash := "$argon2id$v=99999999999$m=65536,t=3,p=2$2DfbJd0beYSdpFYOK3eh7w$hgKoVCc9I31oXV59YCXyeBHUC7EEZQc2byAYy7M74io"
	match, err := Verify("password", hash)
	assert.NotNil(t, err)

	assert.False(t, match)
	assert.Equal(t, err, errors.New("incompatible version of argon2"))
}

func TestInvalidValues(t *testing.T) {
	hash := "$argon2id$TOO_MANY,$v=19$m=65536,t=3,p=2$2DfbJd0beYSdpFYOK3eh7w$hgKoVCc9I31oXV59YCXyeBHUC7EEZQc2byAYy7M74io"
	match, err := Verify("password", hash)
	assert.NotNil(t, err)

	assert.False(t, match)
	assert.Equal(t, err, errors.New("the encoded hash is not in the correct format"))
}
