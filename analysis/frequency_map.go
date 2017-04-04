package analysis

import ()

type CharacterFrequency map[byte]int

func GetCharacterFrequency(b []byte) CharacterFrequency {
	frequencies := map[byte]int{}

	for i := 0; i < len(b); i++ {
		frequencies[b[i]]++
	}

	return CharacterFrequency(frequencies)
}

func (f CharacterFrequency) GetCharacterFrequency(b byte) int {
	v, ok := f[b]
	if !ok {
		return 0
	}
	return v
}

func (f CharacterFrequency) GetTotalCharacters() int {
	total := 0

	for _, v := range f {
		total = total + v
	}

	return total
}

func (f CharacterFrequency) GetCharacterRatios() CharacterRatios {
	total := f.GetTotalCharacters()

	ratios := map[byte]float64{}

	for k, v := range f {
		ratios[k] = float64(v) / float64(total)
	}

	return CharacterRatios(ratios)
}

type CharacterRatios map[byte]float64

func (f CharacterRatios) GetCharacterRatio(b byte) float64 {
	v, ok := f[b]
	if !ok {
		return float64(0)
	}
	return v
}

var EnglishLetterFrequences = map[byte]float64{
	byte('a'): 0.11602,
	byte('b'): 0.04702,
	byte('c'): 0.03511,
	byte('d'): 0.0267,
	byte('e'): 0.02007,
	byte('f'): 0.03779,
	byte('g'): 0.0195,
	byte('h'): 0.07232,
	byte('i'): 0.06286,
	byte('j'): 0.00597,
	byte('k'): 0.0059,
	byte('l'): 0.02705,
	byte('m'): 0.04383,
	byte('n'): 0.02365,
	byte('o'): 0.06264,
	byte('p'): 0.02545,
	byte('q'): 0.00173,
	byte('r'): 0.01653,
	byte('s'): 0.07755,
	byte('t'): 0.16671,
	byte('u'): 0.01487,
	byte('v'): 0.00649,
	byte('w'): 0.06753,
	byte('x'): 0.00017,
	byte('y'): 0.0162,
	byte('z'): 0.00034,

	//just use the same probabilities so we don't tamper with the input
	byte('A'): 0.11602,
	byte('B'): 0.04702,
	byte('C'): 0.03511,
	byte('D'): 0.0267,
	byte('E'): 0.02007,
	byte('F'): 0.03779,
	byte('G'): 0.0195,
	byte('H'): 0.07232,
	byte('I'): 0.06286,
	byte('J'): 0.00597,
	byte('K'): 0.0059,
	byte('L'): 0.02705,
	byte('M'): 0.04383,
	byte('N'): 0.02365,
	byte('O'): 0.06264,
	byte('P'): 0.02545,
	byte('Q'): 0.00173,
	byte('R'): 0.01653,
	byte('S'): 0.07755,
	byte('T'): 0.16671,
	byte('U'): 0.01487,
	byte('V'): 0.00649,
	byte('W'): 0.06753,
	byte('X'): 0.00017,
	byte('Y'): 0.0162,
	byte('Z'): 0.00034,
}
