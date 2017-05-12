package set01

import (
	"cryptobuds/model"
	"cryptobuds/xor"
	"log"
	"testing"
)

func Test_Fixed_XOR(t *testing.T) {
	s1 := model.NewHexString("1c0111001f010100061a024b53535009181c")
	s2 := model.NewHexString("686974207468652062756c6c277320657965")

	s3, err := xor.XORBlockCipher{}.Apply(s1, s2)
	if err != nil {
		t.Fatal(err)
	}

	if string(s3) != "746865206b696420646f6e277420706c6179" {
		t.Fatal("challenge 02 - INCORRECT!")
	} else {
		log.Println("challenge 02 - correct!")
	}
}
