package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomBytes(length int) []byte {
	b := make([]byte, length)

	_, err := rand.Read(b) //math.Intn sucks

	if err != nil { //i don't want to deal with catching and propagating an err up the stack
		panic(err)
	}

	return b
}

func GenerateRandomInteger(min, max int) int {
	big, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil { //i don't want to deal with catching and propagating an err up the stack
		panic(err)
	}

	n := big.Int64()

	if n < int64(min) {
		return GenerateRandomInteger(min, max)
	} else {
		//do nothing
	}

	return int(n)
}
