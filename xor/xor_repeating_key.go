package xor

import (
	"cryptobuds/model"
	"encoding/hex"
)

type XORRepeatingKeyCipher struct {
}

func (c XORRepeatingKeyCipher) Decrypt(cipherText model.HexString, key string) (string, error) {
	cipherTextBytes, err := hex.DecodeString(string(cipherText))
	if err != nil {
		return "", err
	}

	plainTextHex := c.repeatingKeyXOR(string(cipherTextBytes), key)

	decodedPlainText, err := hex.DecodeString(string(plainTextHex))
	if err != nil {
		return "", err
	}

	return string(decodedPlainText), nil
}

func (c XORRepeatingKeyCipher) Encrypt(plainText string, key string) model.HexString {
	return c.repeatingKeyXOR(plainText, key)
}

func (c XORRepeatingKeyCipher) repeatingKeyXOR(in string, key string) model.HexString {
	keyBytes := []byte(key)
	index := 0

	out := []byte{}

	for _, inByte := range []byte(in) {
		outByte := inByte ^ keyBytes[index]

		out = append(out, outByte)

		if index == len(keyBytes)-1 {
			index = 0
		} else {
			index = index + 1
		}
	}

	return model.HexString(hex.EncodeToString(out))
}
