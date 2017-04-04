package analysis

import (
	"log"
	"testing"
)

func Test_Exact_Match(t *testing.T) {
	expected := GetCharacterFrequency([]byte("from the depths of the sea")).GetCharacterRatios()
	actual := GetCharacterFrequency([]byte("from the depths of the sea")).GetCharacterRatios()

	score := ChiSquaredComparer{}.Compare(expected, actual)

	if score.OverallScore != 0 {
		t.Fatal("0 expected")
	} else {
		log.Print("score is 0")
	}
}

func Test_With_Space(t *testing.T) {
	expected := CharacterRatios(EnglishLetterFrequences)
	actual := GetCharacterFrequency([]byte("from the depths of the sea")).GetCharacterRatios()

	score := ChiSquaredComparer{}.Compare(expected, actual)

	log.Println(score.OverallScore)
}

func Test_Only_Letters(t *testing.T) {
	expected := CharacterRatios(EnglishLetterFrequences)
	actual := GetCharacterFrequency([]byte("fromthedepthsofthesea")).GetCharacterRatios()

	score := ChiSquaredComparer{}.Compare(expected, actual)

	log.Println(score.OverallScore)
}

func Test_Complete_Nonsense(t *testing.T) {
	expected := CharacterRatios(EnglishLetterFrequences)
	actual := GetCharacterFrequency([]byte{1, 25, 16, 204, 56, 11, 11, 113, 17, 78}).GetCharacterRatios()

	score := ChiSquaredComparer{}.Compare(expected, actual)

	log.Println(score.OverallScore)
}
