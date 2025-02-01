package jwt

import (
	"testing"

	"github.com/golang-module/carbon"
	"github.com/judegiordano/gogetem/pkg/nanoid"
	"github.com/stretchr/testify/assert"
)

type MyCLaims struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	Aud   string `json:"aud"`
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
	Iss   string `json:"iss"`
	Sub   string `json:"sub"`
}

func claims(expires int64) Claims {
	n, _ := nanoid.New()
	return Claims{
		"id":    n,
		"admin": false,
		"sub":   "terminal",
		"iss":   "testing",
		"aud":   "stdout",
		"exp":   expires,
		"iat":   carbon.Now().Carbon2Time().Unix(),
	}
}

func TestSign(t *testing.T) {
	claims := claims(carbon.Now().AddMinutes(10).Carbon2Time().Unix())
	token, err := Sign(claims)
	assert.Nil(t, err)
	assert.True(t, len(token) > 1)
}

func TestVerifyExpired(t *testing.T) {
	c := claims(carbon.Now().SubDays(1).Carbon2Time().Unix())
	token, err := Sign(c)
	assert.Nil(t, err)

	claims, err := Verify[MyCLaims](token)
	assert.NotNil(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, err.Error(), "token has invalid claims: token is expired")
}

func TestVerifyValid(t *testing.T) {
	c := claims(carbon.Now().AddDay().Carbon2Time().Unix())
	token, err := Sign(c)
	assert.Nil(t, err)

	claims, err := Verify[MyCLaims](token)
	assert.Nil(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, claims.ID, c["id"])
}
