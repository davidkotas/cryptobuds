package xor

import (
	//"log"
	"cryptobuds/analysis"
	"math"
)

type XORSingleByteSolver struct {
}

func (this XORSingleByteSolver) Solve(cipherText []byte) (byte, float64) {
	c := XORSingleByteCipher{}

	chiSquared := analysis.ChiSquaredComparer{}

	min := math.MaxFloat64
	char := byte(' ')

	for i := 0; i < 256; i++ {
		o := c.Apply(cipherText, byte(i))

		actual := analysis.GetCharacterFrequency(o).GetCharacterRatios()

		withSpace := map[byte]float64{}
		for k, v := range analysis.EnglishLetterFrequences {
			withSpace[k] = v
		}
		withSpace[byte(' ')] = 0.20
		withSpace[byte('\'')] = 0.05

		result := chiSquared.Compare(withSpace, actual)

		if result.OverallScore < min {
			if min == math.MaxFloat64 {
				//log.Printf("start: %f, %q", result.OverallScore, byte(i))
			} else {
				//log.Printf("%f -> %f, %q -> %q", min, result.OverallScore, char, byte(i))
			}
			min = result.OverallScore
			char = byte(i)
		}
	}

	return char, min
}
