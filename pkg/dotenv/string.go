package dotenv

import "os"

func String(key string) *string {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		return nil
	}
	return &value
}
