package dotenv

import (
	"os"
	"strconv"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func Bool(key string) *bool {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error("ENV VAR NOT FOUND", key)
		return nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		logger.Error("ENV PARSING BOOL", key)
		return nil
	}
	return &b
}
