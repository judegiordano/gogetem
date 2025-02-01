package jwt

import jwt5 "github.com/golang-jwt/jwt/v5"

func Sign(claims Claims) (string, error) {
	token := jwt5.NewWithClaims(jwt5.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
