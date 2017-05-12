package set01

import (
	"cryptobuds/analysis"
	"cryptobuds/model"
	"cryptobuds/xor"
	"log"
	"testing"
)

func Test_Single_byte_XOR_cipher(t *testing.T) {
	h := model.NewHexString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	c := xor.XORSingleByteCipher{}

	chiSquared := analysis.ChiSquaredComparer{}

	for i := 0; i < 256; i++ {
		b, err := h.GetBytes()
		if err != nil {
			t.Fatal(err)
		}

		o := c.Apply(b, byte(i))

		actual := analysis.GetCharacterFrequency(o).GetCharacterRatios()

		withSpace := map[byte]float64{}
		for k, v := range analysis.EnglishLetterFrequences {
			withSpace[k] = v
		}
		withSpace[byte(' ')] = 0.20
		withSpace[byte('\'')] = 0.05

		result := chiSquared.Compare(withSpace, actual)

		if result.OverallScore < 1 {
			if i != 88 {
				t.Fatal("challenge 03 - INCORRECT!")
			} else {
				log.Println("challenge 03 - correct!")
			}
		}
	}
}
