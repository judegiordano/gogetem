package jwt

import (
	"errors"

	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
)

func Verify[CustomClaims interface{}](token string) (*CustomClaims, error) {
	parsed, err := jwt5.Parse(token, func(token *jwt5.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	m, ok := parsed.Claims.(jwt5.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}
	var claims CustomClaims
	if err := mapstructure.Decode(m, &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}
