package dotenv

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func normalizeKey(k string) string {
	return strings.TrimSpace(strings.ToUpper(k))
}

func String(key string) *string {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		return nil
	}
	return &value
}

func Int(key string) *int {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		return nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}
	return &i
}

func Bool(key string) *bool {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		return nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return &b
}

func init() {
	if err := godotenv.Load(); err != nil {
		// TODO: maybe log
	}
}
