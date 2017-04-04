package set02

import (
	"cryptopals/aes"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_An_ECB_CBC_detection_oracle(t *testing.T) {
	blackbox := aes.AesBlockBox{}
	oracle := aes.EcbOracle{}

	//https://en.wikipedia.org/wiki/Known-plaintext_attack

	crib := []byte{}

	for i := 0; i < 2560; i++ {
		crib = append(crib, 'a')
	}

	for i := 0; i < 20; i++ {
		mode, encrypted, err := blackbox.Encrypt(string(crib))
		if err != nil {
			t.Fatal(err)
		}

		isEcb := oracle.IsEcb(encrypted)

		if isEcb != (mode == aes.ECB) {
			log.Printf("[Expected] [%s], [Actual] [%s]", mode, aes.ECB)
			t.Fatal("challenge 11 - incorrect!")
		}
	}
	log.Println("challenge 11 - correct!")
}
