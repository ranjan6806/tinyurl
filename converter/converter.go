package converter

import (
	"crypto/sha256"
	"encoding/base64"
)

type Converter interface {
	Encode(url string) string
}

type HashConverter struct {
}

func NewHashConverter() *HashConverter {
	return &HashConverter{}
}

func (hc *HashConverter) Encode(url string) string {
	hashedURL := sha256.Sum256([]byte(url))
	encodedURL := base64.URLEncoding.EncodeToString(hashedURL[:10])
	return encodedURL
}
