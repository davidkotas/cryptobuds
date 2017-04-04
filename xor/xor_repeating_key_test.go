package xor

import (
	"log"
	"testing"
)

func Test_Round_Trip(t *testing.T) {
	cipher := XORRepeatingKeyCipher{}

	in := "So won't the real Slim Shady please stand up, Please stand up, please stand up?"
	key := "SLIM"

	cipherText := cipher.Encrypt(in, key)

	plain, err := cipher.Decrypt(cipherText, key)
	if err != nil {
		t.Fatal(err)
	}

	if plain != in {
		t.Fatal("input not recovered.")
	} else {
		log.Println("input recovered!")
	}
}
