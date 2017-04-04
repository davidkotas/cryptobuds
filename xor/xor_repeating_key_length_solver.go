package xor

import (
	"cryptopals/analysis"
	"cryptopals/model"
)

type XORRepeatingKeyLengthCalculator struct {
}

func (this XORRepeatingKeyLengthCalculator) WithFirstTwoBlocks(cipherText model.HexString, length int) (float64, error) {
	chunks, err := this.chunk(cipherText, length)
	if err != nil {
		return 0.0, err
	}

	x := chunks[0]
	y := chunks[1]

	distance, err := analysis.HammingDistanceCalculator{}.CalculateBytes(x, y)
	if err != nil {
		return 0.0, err
	}

	normalized := float64(distance) / float64(length)

	return normalized, nil
}

func (this XORRepeatingKeyLengthCalculator) WithProgressiveTwoBlocks(cipherText model.HexString, length int) (float64, error) {
	distances := []float64{}

	chunks, err := this.chunk(cipherText, length)
	if err != nil {
		return 0.0, err
	}

	for i := 0; i < (len(chunks)-1)-1; i++ {
		x := chunks[i]
		y := chunks[i+1]

		distance, err := analysis.HammingDistanceCalculator{}.CalculateBytes(x, y)
		if err != nil {
			return 0.0, err
		}

		normalized := float64(distance) / float64(length)

		distances = append(distances, normalized)
	}

	sum := 0.0
	for _, distance := range distances {
		sum += distance
	}

	return sum / float64(len(distances)), nil
}

func (this XORRepeatingKeyLengthCalculator) WithAllBlocks(cipherText model.HexString, length int) (float64, error) {
	distances := []float64{}

	chunks, err := this.chunk(cipherText, length)
	if err != nil {
		return 0.0, err
	}

	for i := 0; i < len(chunks)-1; i++ {
		for j := 0; j < len(chunks)-1; j++ {
			if i != j {
				x := chunks[i]
				y := chunks[j]

				distance, err := analysis.HammingDistanceCalculator{}.CalculateBytes(x, y)
				if err != nil {
					return 0.0, err
				}

				normalized := float64(distance) / float64(length)

				distances = append(distances, normalized)
			}
		}
	}

	sum := 0.0
	for _, distance := range distances {
		sum += distance
	}

	return sum / float64(len(distances)), nil
}

func (this XORRepeatingKeyLengthCalculator) chunk(cipherText model.HexString, length int) (model.ByteMatrix, error) {
	cipherTextHexBytes, err := cipherText.GetBytes()
	if err != nil {
		return model.ByteMatrix([][]byte{}), err
	}

	array := model.ByteArray(cipherTextHexBytes.GetBytes())

	chunks := array.Chunk(length)

	return chunks, nil
}
