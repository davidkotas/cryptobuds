package xor

import ()

type XORSingleByteCipher struct {
}

func (c XORSingleByteCipher) Apply(input []byte, key byte) []byte {
	output := []byte{}

	for _, b := range input {
		output = append(output, b^key)
	}

	return output
}
