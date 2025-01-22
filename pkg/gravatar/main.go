package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateHash(email string) string {
	norm := strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(norm))
	return hex.EncodeToString(hash[:])
}

func Url(email string) string {
	hash := GenerateHash(email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro", hash)
}
