package analysis

import (
	"log"
	"testing"
)

func Test_Hamming(t *testing.T) {
	calc := HammingDistanceCalculator{}

	distance, err := calc.CalculateString("this is a test", "wokka wokka!!!")
	if err != nil {
		t.Fatal(err)
	}

	if distance != 37 {
		t.Fatal("hamming test failed!")
	} else {
		log.Println("hamming test passed.")
	}
}
