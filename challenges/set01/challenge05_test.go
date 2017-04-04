package set01

import (
	"cryptopals/xor"
	"log"
	"testing"
)

func Test_Implement_repeating_key_XOR_01(t *testing.T) {
	in := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	cipherText := xor.XORRepeatingKeyCipher{}.Encrypt(in, "ICE")

	if string(cipherText) !=
		"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		t.Fatal("challenge 05 - INCORRECT!")
	} else {
		log.Println("challenge 05 - correct!")
	}
}
