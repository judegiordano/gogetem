package gravatar

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	email := "mail@mail.com"
	hash := generateHash(email)
	assert.Equal(t, hash, "7905d373cfab2e0fda04b9e7acc8c879")
	url := Url(email)
	assert.Equal(t, url, fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro", hash))
}
