package aes

import (
	"cryptopals/model"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_AesEcb_Encrypt_Decrypt(t *testing.T) {
	key := "WE RUST IN PEACE"

	ecb := AesEcb{}

	encrypted, err := ecb.Encrypt(plain, key)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := ecb.DecryptBytes(encrypted, key)
	if err != nil {
		t.Fatal(err)
	}

	if !model.Equals(decrypted, []byte(plain)) {
		t.Fatal("plain and decrypted do not match.")
	} else {
		log.Println("Test_AesCbc_Encrypt_Decrypt pass!")
	}
}
