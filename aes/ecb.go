package aes

import (
	"crypto/aes"
	"cryptobuds/model"
	"cryptobuds/padding"
)

type AesEcb struct {
}

func (this AesEcb) Encrypt(plain, key string) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	plainBlocks := model.ByteArray([]byte(plain)).Chunk(cipher.BlockSize())

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

		cipherBlock := make([]byte, cipher.BlockSize(), cipher.BlockSize())
		cipher.Encrypt(cipherBlock, plainBlock)

		cipherBlocks = append(cipherBlocks, cipherBlock)
	}

	final := []byte{}

	for _, cipherBlock := range cipherBlocks {
		final = append(final, cipherBlock...)
	}

	return final, nil
}

func (this AesEcb) DecryptBytes(cipherText []byte, key string) ([]byte, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	cipherBlocks := model.ByteArray(cipherText).Chunk(cipher.BlockSize())

	plainBlocks := [][]byte{}

	for i, cipherBlock := range cipherBlocks {
		plainBlock := make([]byte, cipher.BlockSize(), cipher.BlockSize())
		cipher.Decrypt(plainBlock, cipherBlock)

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
	}

	final := []byte{}

	for _, plainBlock := range plainBlocks {
		final = append(final, plainBlock...)
	}

	return final, nil
}

func (this AesEcb) Decrypt(hexString model.HexString, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	cipherBytes, err := hexString.GetBytes()
	if err != nil {
		return nil, err
	}

	chunks := model.ByteArray([]byte(cipherBytes)).Chunk(block.BlockSize())

	plain := []byte{}

	for _, chunk := range chunks {
		dst := make([]byte, block.BlockSize(), block.BlockSize())
		block.Decrypt(dst, chunk)
		plain = append(plain, dst...)
	}

	return plain, nil
}
