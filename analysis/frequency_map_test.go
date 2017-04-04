package analysis

import (
	"log"
	"testing"
)

func Test_Correct_Frequency(t *testing.T) {
	s1 := "she sells sea shells by the sea shore"

	frequencies := GetCharacterFrequency([]byte(s1))

	s := frequencies.GetCharacterFrequency(byte('s'))
	if s != 8 {
		t.Fatal("wrong frequency for s")
	} else {
		log.Println("frequency ok for s")
	}
}

func Test_Correct_Ratios(t *testing.T) {
	s1 := "bebop"

	frequencies := GetCharacterFrequency([]byte(s1))

	ratios := frequencies.GetCharacterRatios()

	if ratios.GetCharacterRatio(byte('e')) != 0.200000 {
		t.Fatal("wrong ratio for e")
	} else {
		log.Println("ratio ok for e")
	}
	if ratios.GetCharacterRatio(byte('o')) != 0.200000 {
		t.Fatal("wrong ratio for o")
	} else {
		log.Println("ratio ok for o")
	}
	if ratios.GetCharacterRatio(byte('p')) != 0.200000 {
		t.Fatal("wrong ratio for p")
	} else {
		log.Println("ratio ok for p")
	}
	if ratios.GetCharacterRatio(byte('b')) != 0.400000 {
		t.Fatal("wrong ratio for b")
	} else {
		log.Println("ratio ok for b")
	}
}
