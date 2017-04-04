package util

func GenerateCrib(length int, character byte) []byte {
	crib := []byte{}

	for i := 0; i < length; i++ {
		crib = append(crib, character)
	}

	return crib
}
