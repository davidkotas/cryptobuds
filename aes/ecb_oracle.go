package aes

import (
	"cryptopals/model"
)

type EcbOracle struct {
}

func (this EcbOracle) IsEcb(cipher []byte) bool {
	chunks := (model.ByteArray(cipher)).Chunk(16) //we already know it's 16 bytes, is this assumption valid for this exercise?

	for i := 0; i < len(chunks); i++ {
		for j := 0; j < len(chunks); j++ {
			if i == j {
				continue
			}
			e := model.Equals(chunks[i], chunks[j])
			if e {
				return true //if any blocks are identical, ASSUME ECB.  what's the probability any other mode will give two identical blocks?
			}
		}
	}

	return false
}
