package padding

import (
	"log"
	"testing"
)

func Test_AddPadding(t *testing.T) {
	message := "YELLOW SUBMARINE"
	block := []byte(message)

	padded, _ := PKCS7{}.AddPadding(block, 20)

	for i := 16; i < 20; i++ {
		if padded[i] != byte(4) {
			t.Fatal(padded[i])
		}
	}
}

func Test_RemovePadding_Success(t *testing.T) {
	message := "YELLOW SUBMARINE"
	block := []byte(message)

	padder := PKCS7{}

	padded, _ := padder.AddPadding(block, 20)

	unpadded, _ := padder.RemovePadding(padded, 20)

	if string(unpadded) != "YELLOW SUBMARINE" {
		t.Fatal("unpadding failed")
	} else {
		log.Println(string(unpadded))
	}
}

func Test_RemovePadding_Failed(t *testing.T) {
	message := "YELLOW SUBMARINE"
	block := []byte(message)

	block = append(block, byte(1))
	block = append(block, byte(2))
	block = append(block, byte(3))
	block = append(block, byte(4))

	padder := PKCS7{}

	_, err := padder.RemovePadding(block, 20)
	if err == nil {
		t.Fatal("expecting error")
	} else {
		log.Println(err)
	}
}
