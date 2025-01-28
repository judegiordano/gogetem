package dotenv

import (
	"strings"

	"github.com/joho/godotenv"
)

func normalizeKey(k string) string {
	return strings.TrimSpace(strings.ToUpper(k))
}

func init() {
	if err := godotenv.Load(); err != nil {
		// TODO: maybe log
	}
}
