package xor

import (
	"log"
	"testing"
)

func Test_Single_Byte_Solver(t *testing.T) {
	testSingleByteSolver(t, "this is the plain text.", byte('b'))
	testSingleByteSolver(t, "this is the plain text.", byte('c'))
	testSingleByteSolver(t, "this is the plain text.", byte('d'))
	testSingleByteSolver(t, "this is the plain text.", byte(' '))
	testSingleByteSolver(t, "this is the plain text.", byte('g'))
	testSingleByteSolver(t, "this is the plain text.", byte('f'))
	testSingleByteSolver(t, "So won't the real Slim Shady please stand up, Please stand up, please stand up?", 'S')
}

func testSingleByteSolver(t *testing.T, plainText string, key byte) {
	encrypted := XORSingleByteCipher{}.Apply([]byte(plainText), key)

	solved, _ := XORSingleByteSolver{}.Solve(encrypted)

	if solved != key {
		t.Fatal("key not recovered.")
	} else {
		log.Println("key recovered!")
	}
}
