package dotenv

import (
	"os"
	"strconv"
)

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
