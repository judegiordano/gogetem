package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func generateHash(email string) string {
	norm := strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(norm))
	return hex.EncodeToString(hash[:])
}

func Url(email string) string {
	hash := generateHash(email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro", hash)
}
