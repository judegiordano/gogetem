package dotenv

import (
	"os"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func String(key string) *string {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error("ENV VAR NOT FOUND", key)
		return nil
	}
	return &value
}
