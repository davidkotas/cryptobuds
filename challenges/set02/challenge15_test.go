package set02

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_PKCS7_padding_validation(t *testing.T) {
	//see tests in pkcs7_test.go
	log.Println("challenge 15 - correct!")
}
