package set02

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_Implement_PKCS7_padding(t *testing.T) {
	//see tests in pkcs7_test.go
	log.Println("challenge 09 - correct!")
}
