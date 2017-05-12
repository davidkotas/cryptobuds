package aes

import (
	"crypto/aes"
	"cryptobuds/model"
	"cryptobuds/padding"
	"cryptobuds/xor"
	"errors"
)

type AesCbc struct {
}

func (this AesCbc) Encrypt(plain, key, iv string) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if len(iv) != cipher.BlockSize() {
		return nil, errors.New("incorrect iv size")
	}

	plainBlocks := model.ByteArray([]byte([]byte(plain))).Chunk(cipher.BlockSize())

	xorBlockCipher := xor.XORBlockCipher{}

	//hence the name initialization vector
	lastCipherBlock := []byte(iv)

	cipherBlocks := [][]byte{}

	for i, plainBlock := range plainBlocks {
		if i == len(plainBlocks)-1 {
			if len(plainBlock) != cipher.BlockSize() {
				padded, err := padding.PKCS7{}.AddPadding(plainBlock, cipher.BlockSize())
				if err != nil {
					return nil, err
				}
				plainBlock = padded
			}
		}

		xorBlock, err := xorBlockCipher.ApplyToBytes(lastCipherBlock, plainBlock)
		if err != nil {
			return nil, err
		}

		cipherBlock := make([]byte, cipher.BlockSize(), cipher.BlockSize())
		cipher.Encrypt(cipherBlock, xorBlock)

		cipherBlocks = append(cipherBlocks, cipherBlock)

		lastCipherBlock = cipherBlock
	}

	final := []byte{}

	for _, cipherBlock := range cipherBlocks {
		final = append(final, cipherBlock...)
	}

	return final, nil
}

func (this AesCbc) Decrypt(cipherText []byte, key, iv string) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	if len(iv) != cipher.BlockSize() {
		return nil, errors.New("incorrect iv size")
	}

	cipherBlocks := model.ByteArray(cipherText).Chunk(cipher.BlockSize())

	xorBlockCipher := xor.XORBlockCipher{}

	lastCipherBlock := []byte(iv)

	plainBlocks := [][]byte{}

	for i, cipherBlock := range cipherBlocks {
		xorBlock := make([]byte, cipher.BlockSize(), cipher.BlockSize())
		cipher.Decrypt(xorBlock, cipherBlock)

		plainBlock, err := xorBlockCipher.ApplyToBytes(lastCipherBlock, xorBlock)
		if err != nil {
			return nil, err
		}

		if i == len(cipherBlocks)-1 {
			//assume that the last cipherblock is padded, and always remove
			//could be a bad assumption...we'll see
			unpadded, err := padding.PKCS7{}.RemovePadding(plainBlock, cipher.BlockSize())
			if err != nil {
				return nil, err
			}
			plainBlock = unpadded
		}

		plainBlocks = append(plainBlocks, plainBlock)

		lastCipherBlock = cipherBlock
	}

	final := []byte{}

	for _, plainBlock := range plainBlocks {
		final = append(final, plainBlock...)
	}

	return final, nil
}
