package set01

import (
	"cryptopals/model"
	"cryptopals/xor"
	"encoding/base64"
	"io/ioutil"
	"log"
	"sort"
	"testing"
)

func Test_Break_repeating_key_XOR(t *testing.T) {
	base64Bytes, err := ioutil.ReadFile("challenge06_text.txt")
	if err != nil {
		t.Fatal(err)
	}

	cipherText, err := base64.StdEncoding.DecodeString(string(base64Bytes))
	if err != nil {
		t.Fatal(err)
	}

	cipherTextBytes := []byte(cipherText)

	keySizeGuesses, err := guessKeySize(cipherTextBytes, 3)
	if err != nil {
		t.Fatal(err)
	}

	//after much consternation, we know the first guess will be correct
	guess := keySizeGuesses[0]

	key := solve(cipherTextBytes, guess.KeyLength)

	if key != "Terminator X: Bring the noise" {
		t.Fatal("challenge 06 - INCORRECT!")
	} else {
		log.Println("challenge 06 - correct!")
	}
}

type KeySizeGuess struct {
	KeyLength int
	Distance  float64
}

type byDistance []KeySizeGuess

func (x byDistance) Len() int           { return len(x) }
func (x byDistance) Less(i, j int) bool { return x[i].Distance < x[j].Distance }
func (x byDistance) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func guessKeySize(cipherBytes []byte, take int) ([]KeySizeGuess, error) {
	minKeySize := 2
	maxKeySize := 40

	guesses := []KeySizeGuess{}

	calculator := xor.XORRepeatingKeyLengthCalculator{}

	for keyLength := minKeySize; keyLength < maxKeySize+1; keyLength++ {
		score, err := calculator.WithAllBlocks((model.HexBytes(cipherBytes)).GetString(), keyLength)
		if err != nil {
			return nil, err
		}

		guesses = append(guesses, KeySizeGuess{
			KeyLength: keyLength,
			Distance:  score,
		})
	}

	sort.Sort(byDistance(guesses))

	return guesses[0:take], nil
}

func solve(cipherBytes []byte, keySize int) string {
	ba := model.ByteArray(cipherBytes)

	blocks := ba.Chunk(keySize)

	transposed := blocks.Transpose()

	solver := xor.XORSingleByteSolver{}

	key := ""
	for _, block := range transposed {
		char, _ := solver.Solve(block)
		key = key + string(char)
	}
	return key
}
