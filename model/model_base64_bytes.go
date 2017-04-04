package model

import (
	"encoding/base64"
	"encoding/hex"
)

type Base64Bytes []byte

func (this Base64Bytes) GetHexBytes() (HexBytes, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(this))
	if err != nil {
		return HexBytes([]byte{}), err
	}

	hexEncoded := hex.EncodeToString(decoded)

	return HexBytes([]byte(hexEncoded)), nil
}
