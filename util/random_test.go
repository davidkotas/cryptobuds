package util

import (
	"log"
	"testing"
)

func Test_GenerateRandomInteger(t *testing.T) {
	i := GenerateRandomInteger(25, 1001)

	if i > 1001 {
		t.Fatal("wrong random")
	} else {
		log.Println(i)
	}
}
