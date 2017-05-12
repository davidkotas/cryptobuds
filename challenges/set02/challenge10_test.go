package set02

import (
	"cryptobuds/aes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_Implement_CBC_mode(t *testing.T) {
	base64Bytes, err := ioutil.ReadFile("challenge10_text.txt")
	if err != nil {
		t.Fatal(err)
	}

	cipherText, err := base64.StdEncoding.DecodeString(string(base64Bytes))
	if err != nil {
		t.Fatal(err)
	}

	plainText, err := aes.AesCbc{}.Decrypt(cipherText, "YELLOW SUBMARINE", "0000000000000000")
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(plainText), "It controls my mouth and I begin") {
		log.Println("challenge 10 - correct!")
	} else {
		t.Fatal("challenge 10 - INCORRECT!")
	}
}
