package dotenv

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		logger.Debug("no .env present: [%v]", err)
	}
}

func String(key string) *string {
	normalized := normalizeKey(key)
	println("normalized: ", normalized)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Warn(fmt.Sprintf(".env %v not set", normalized))
		return nil
	}
	return &value
}

func normalizeKey(k string) string {
	return strings.TrimSpace(strings.ToUpper(k))
}
