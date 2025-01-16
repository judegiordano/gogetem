package dotenv

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/judegiordano/gogetem/pkg/logger"
)

func normalizeKey(k string) string {
	return strings.TrimSpace(strings.ToUpper(k))
}

func String(key string) *string {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error(fmt.Sprintf(".env %v not set", normalized))
		return nil
	}
	return &value
}

func Int(key string) *int {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error(fmt.Sprintf(".env %v not set", normalized))
		return nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot parse %v as int %v", value, err))
		return nil
	}
	return &i
}

func Bool(key string) *bool {
	normalized := normalizeKey(key)
	value, found := os.LookupEnv(normalized)
	if !found {
		logger.Error(".env %v not set", normalized)
		return nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		logger.Error("cannot parse %v as bool %v", value, err)
		return nil
	}
	return &b
}

func init() {
	logger.Debug("loading .env...")
	if err := godotenv.Load(); err != nil {
		logger.Warn("no .env present", err)
	}
}
