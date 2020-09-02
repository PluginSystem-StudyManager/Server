package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateToken() string {
	const tokenLength = 64
	b := make([]byte, tokenLength)
	_, _ = rand.Read(b)
	token := hex.EncodeToString(b)
	return token
}
