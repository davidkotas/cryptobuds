package padding

import (
	"errors"
	"log"
)

type PKCS7 struct {
}

func (this PKCS7) AddPadding(block []byte, length int) ([]byte, error) {
	if len(block) > length {
		return nil, errors.New("block length > length")
	}

	if len(block) == length {
		return block, nil
	}

	padded := []byte{}

	for i := 0; i < len(block); i++ {
		padded = append(padded, block[i])
	} //don't tamper with the input

	required := length - len(block)

	for i := 0; i < required; i++ {
		padded = append(padded, byte(required))
	}

	return padded, nil
}

func (this PKCS7) RemovePadding(block []byte, length int) ([]byte, error) {
	if len(block) != length {
		return nil, errors.New("unexpected block length.")
	}

	padEnd := len(block) - 1

	lastByte := block[padEnd]

	padStart := len(block) - int(lastByte)

	padding := block[padStart : padEnd+1]

	//validate padding
	for i := 0; i < len(padding); i++ {
		if padding[i] != lastByte {
			log.Printf("[%d != %d]", padding[i], lastByte)
			return nil, errors.New("invalid padding.")
		} else {
			//do nothing
		}
	}

	unpadded := []byte{}
	for i := 0; i < padStart; i++ {
		unpadded = append(unpadded, block[i])
	}

	return unpadded, nil
}
