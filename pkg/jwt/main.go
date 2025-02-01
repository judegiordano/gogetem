package jwt

import (
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/judegiordano/gogetem/pkg/dotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
)

type Claims = jwt5.MapClaims

var key []byte

func init() {
	secret := dotenv.String("JWT_SECRET")
	if secret == nil {
		logger.Fatal("[JWT]", "JWT_SECRET is required")
	}
	key = []byte(*secret)
}
