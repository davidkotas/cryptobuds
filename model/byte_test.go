package model

import (
	"log"
	"testing"
)

func Test_Chunk(t *testing.T) {
	bytes := []byte{'s', 'l', 'i', 'm', 's', 'h', 'a', 'd', 'y', ' ', 'r', 'e', 'a', 'l'}

	chunked := ByteArray(bytes).Chunk(4)

	assertCharacterAtPosition(t, chunked, 0, 0, byte('s'))
	assertCharacterAtPosition(t, chunked, 0, 1, byte('l'))
	assertCharacterAtPosition(t, chunked, 0, 2, byte('i'))
	assertCharacterAtPosition(t, chunked, 0, 3, byte('m'))

	assertCharacterAtPosition(t, chunked, 1, 0, byte('s'))
	assertCharacterAtPosition(t, chunked, 1, 1, byte('h'))
	assertCharacterAtPosition(t, chunked, 1, 2, byte('a'))
	assertCharacterAtPosition(t, chunked, 1, 3, byte('d'))

	assertCharacterAtPosition(t, chunked, 2, 0, byte('y'))
	assertCharacterAtPosition(t, chunked, 2, 1, byte(' '))
	assertCharacterAtPosition(t, chunked, 2, 2, byte('r'))
	assertCharacterAtPosition(t, chunked, 2, 3, byte('e'))

	assertCharacterAtPosition(t, chunked, 3, 0, byte('a'))
	assertCharacterAtPosition(t, chunked, 3, 1, byte('l'))
}

func Test_Transpose(t *testing.T) {
	bytes := []byte{'s', 'l', 'i', 'm', 's', 'h', 'a', 'd', 'y', ' ', 'r', 'e', 'a', 'l'}

	chunked := ByteArray(bytes).Chunk(4)

	transposed := chunked.Transpose()

	assertCharacterAtPosition(t, transposed, 0, 0, byte('s'))
	assertCharacterAtPosition(t, transposed, 0, 1, byte('s'))
	assertCharacterAtPosition(t, transposed, 0, 2, byte('y'))
	assertCharacterAtPosition(t, transposed, 0, 3, byte('a'))

	assertCharacterAtPosition(t, transposed, 1, 0, byte('l'))
	assertCharacterAtPosition(t, transposed, 1, 1, byte('h'))
	assertCharacterAtPosition(t, transposed, 1, 2, byte(' '))
	assertCharacterAtPosition(t, transposed, 1, 3, byte('l'))

	assertCharacterAtPosition(t, transposed, 2, 0, byte('i'))
	assertCharacterAtPosition(t, transposed, 2, 1, byte('a'))
	assertCharacterAtPosition(t, transposed, 2, 2, byte('r'))

	assertCharacterAtPosition(t, transposed, 3, 0, byte('m'))
	assertCharacterAtPosition(t, transposed, 3, 1, byte('d'))
	assertCharacterAtPosition(t, transposed, 3, 2, byte('e'))
}
func assertCharacterAtPosition(t *testing.T, matrix ByteMatrix, row, column int, expected byte) {
	b := [][]byte(matrix)
	actual := b[row][column]
	if actual != expected {
		t.Fatal("wrong character at position (%d, %d) - %q", row, column, actual)
	} else {
		log.Printf("correct character at position (%d, %d) - %q", row, column, actual)
	}
}
