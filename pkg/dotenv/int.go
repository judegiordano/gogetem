package dotenv

import (
	"os"
	"strconv"

	"github.com/judegiordano/gogetem/pkg/logger"
)

func Int(key string) *int {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error("ENV VAR NOT FOUND", key)
		return nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Error("ENV PARSING INT", key)
		return nil
	}
	return &i
}
