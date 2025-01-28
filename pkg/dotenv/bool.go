package dotenv

import (
	"os"
	"strconv"
)

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
