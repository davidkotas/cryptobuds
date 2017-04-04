package set01

import (
	"cryptopals/model"
	"log"
	"testing"
)

func Test_Convert_hex_to_base64(t *testing.T) {
	h := model.NewHexString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	b64, err := h.GetBase64String()
	if err != nil {
		t.Fatal(err)
	}

	if string(b64) != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal("challenge 01 - INCORRECT!")
	} else {
		log.Println("challenge 01 - correct!")
	}
}
