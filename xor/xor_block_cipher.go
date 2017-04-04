package xor

import (
	"cryptopals/model"
	"errors"
)

type XORBlockCipher struct {
}

func (c XORBlockCipher) Apply(first model.HexString, second model.HexString) (model.HexString, error) {
	if first.Length() != second.Length() {
		return "", errors.New("length mismatch")
	}

	bFirst, err := first.GetBytes()
	if err != nil {
		return "", err
	}

	bSecond, err := second.GetBytes()
	if err != nil {
		return "", err
	}

	n := bFirst.Length()

	b := make([]byte, n)

	for i := 0; i < n; i++ {
		b[i] = bFirst[i] ^ bSecond[i]
	}

	return model.HexBytes(b).GetString(), nil
}

func (c XORBlockCipher) ApplyToBytes(bFirst []byte, bSecond []byte) ([]byte, error) {
	if len(bFirst) != len(bSecond) {
		return nil, errors.New("length mismatch")
	}

	n := len(bFirst)

	b := make([]byte, n)

	for i := 0; i < n; i++ {
		b[i] = bFirst[i] ^ bSecond[i]
	}

	return b, nil
}
