package analysis

import ()

type ChiSquaredCharacterScore struct {
	Character byte
	Score     float64
}

type ChiSquaredComparerResult struct {
	OverallScore float64
	Scores       []ChiSquaredCharacterScore
}

type ChiSquaredComparer struct {
}

func (c ChiSquaredComparer) Compare(expected CharacterRatios, actual CharacterRatios) ChiSquaredComparerResult {
	chiSquared := float64(0.0)

	scores := []ChiSquaredCharacterScore{}

	for actualChar, actualRatio := range actual {
		expectedRatio, ok := expected[actualChar]

		expectedRatioNumerator := float64(0.0)
		expectedRatioDenominator := float64(0.0)

		if !ok {
			expectedRatioNumerator, expectedRatioDenominator = float64(10), float64(1) //an arbitrarily high number, avoid divide by zero
		} else {
			expectedRatioNumerator, expectedRatioDenominator = expectedRatio, expectedRatio
		}

		numerator := (actualRatio - expectedRatioNumerator) * (actualRatio - expectedRatioNumerator)

		term := float64(numerator / expectedRatioDenominator)

		scores = append(scores, ChiSquaredCharacterScore{
			Character: actualChar,
			Score:     term,
		})

		chiSquared = chiSquared + term
	}

	return ChiSquaredComparerResult{
		OverallScore: chiSquared,
		Scores:       scores,
	}
}
