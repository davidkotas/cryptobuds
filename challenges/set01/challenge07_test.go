package set01

import (
	"cryptopals/aes"
	"cryptopals/model"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"log"
	"testing"
)

func Test_AES_in_ECB_mode(t *testing.T) {
	base64Bytes, err := ioutil.ReadFile("challenge07_text.txt")
	if err != nil {
		t.Fatal(err)
	}

	cipherText, err := base64.StdEncoding.DecodeString(string(base64Bytes))
	if err != nil {
		t.Fatal(err)
	}

	hexString := model.NewHexString(hex.EncodeToString(cipherText))

	plainTextBytes, err := aes.AesEcb{}.Decrypt(hexString, "YELLOW SUBMARINE")
	if err != nil {
		t.Fatal(err)
	}

	plainText := ""
	for _, pt := range plainTextBytes {
		plainText = plainText + string(pt)
	}

	log.Println(plainText)
}
