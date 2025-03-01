package dotenv

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
)

func normalizeKey(k string) string {
	return strings.TrimSpace(strings.ToUpper(k))
}

func init() {
	if err := godotenv.Load(); err != nil {
		logger.Warn("DOTENV", "NO .env FILE")
	}
}
