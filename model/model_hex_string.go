package model

import (
	"encoding/base64"
	"encoding/hex"
)

type HexString string

func NewHexString(h string) HexString {
	return HexString(h)
}

func (h HexString) GetBase64String() (Base64String, error) {
	b, err := hex.DecodeString(string(h))
	if err != nil {
		return "", err
	}

	return Base64String(base64.StdEncoding.EncodeToString(b)), nil
}

func (h HexString) GetBytes() (HexBytes, error) {
	b, err := hex.DecodeString(string(h))
	if err != nil {
		return nil, err
	}

	return HexBytes(b), nil
}

func (h HexString) Length() int {
	return len(string(h))
}
