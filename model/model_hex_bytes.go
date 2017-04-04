package model

import (
	"encoding/hex"
)

type HexBytes []byte

func (h HexBytes) GetBytes() []byte {
	return []byte(h)
}

func (h HexBytes) GetString() HexString {
	return HexString(hex.EncodeToString([]byte(h)))
}

func (h HexBytes) Length() int {
	return len([]byte(h))
}
