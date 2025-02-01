package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/judegiordano/gogetem/pkg/logger"
)

type Gravatar struct {
	Email   string
	Size    int
	Default Default
}

type GravatarData struct {
	EmailHash string
	Url       string
	Image     []byte
}

func GenerateHash(email string) string {
	norm := strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(norm))
	return hex.EncodeToString(hash[:])
}

func Url(email string, size string) string {
	hash := GenerateHash(email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro&s=%v", hash, size)
}

func Image(email string, size int) ([]byte, error) {
	s := strconv.Itoa(size)
	url := Url(email, s)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("GRAVATAR REQUEST", err)
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("GRAVATAR BYTE READ", err)
		return nil, err
	}
	return bytes, nil
}

func (g *Gravatar) Generate() (*GravatarData, error) {
	size := strconv.Itoa(g.Size)
	hash := GenerateHash(g.Email)
	var d string
	if len(g.Default) == 0 {
		d = string(Retro)
	} else {
		d = string(g.Default)
	}

	url := fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=%v&s=%v", hash, d, size)
	resp, err := http.Get(url)

	if err != nil {
		logger.Error("GRAVATAR REQUEST", err)
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("GRAVATAR BYTE READ", err)
		return nil, err
	}
	return &GravatarData{EmailHash: hash, Url: url, Image: bytes}, nil
}

type Default string

const (
	MysteryPerson Default = "mp"
	Identicon     Default = "identicon"
	Monsterid     Default = "monsterid"
	Wavatar       Default = "wavatar"
	Retro         Default = "retro"
	Robohash      Default = "robohash"
)
