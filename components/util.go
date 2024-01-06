package components

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func shortUID() string {
	b := make([]byte, 4) // equals 8 characters
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Errorf("rand.Read: %w", err))
	}
	return hex.EncodeToString(b)
}

type CustomElement struct {
	Script string
	Style  string
}
