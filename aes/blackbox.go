package aes

import (
	"cryptopals/util"
)

type AesBlockBox struct {
}

func (this AesBlockBox) Encrypt(plain string) (BlockCipherMode, []byte, error) { //return block cipher mode to check guess of oracle
	prependLength := util.GenerateRandomInteger(5, 10)
	appendLength := util.GenerateRandomInteger(5, 10)

	prependBytes := util.GenerateRandomBytes(prependLength)
	appendBytes := util.GenerateRandomBytes(appendLength)

	working := []byte{}
	working = append(working, prependBytes...)
	working = append(working, []byte(plain)...)
	working = append(working, appendBytes...)

	choice := util.GenerateRandomInteger(1, 100)

	key := util.GenerateRandomBytes(16)

	if choice%2 == 0 {
		cipher := AesEcb{}

		encrypted, err := cipher.Encrypt(plain, string(key))
		if err != nil {
			return ECB, nil, err
		}
		return ECB, encrypted, nil

	} else {
		cipher := AesCbc{}

		iv := util.GenerateRandomBytes(16)

		encrypted, err := cipher.Encrypt(plain, string(key), string(iv))
		if err != nil {
			return CBC, nil, err
		}

		return CBC, encrypted, nil
	}
}
