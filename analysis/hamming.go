package analysis

import (
	"errors"
	"fmt"
)

type HammingDistanceCalculator struct {
}

func (c HammingDistanceCalculator) CalculateString(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length mismatch")
	}

	x := []byte(a)
	y := []byte(b)

	return c.CalculateBytes(x, y)
}

func (c HammingDistanceCalculator) CalculateBytes(x, y []byte) (int, error) {
	if len(x) != len(y) {
		return 0, errors.New("length mismatch")
	}

	difference := 0

	for i := 0; i < len(x); i++ {
		xb := fmt.Sprintf("%08b", int64(x[i]))
		yb := fmt.Sprintf("%08b", int64(y[i]))

		for j := 0; j < len(xb); j++ {
			if xb[j] != yb[j] {
				difference++
			}
		}
	}

	return difference, nil
}
