package model

type ByteArray []byte

func (b ByteArray) Chunk(chunksize int) ByteMatrix {
	chunks := [][]byte{}

	for i := 0; i < len(b); i += chunksize {
		end := i + chunksize
		if end > len(b) {
			end = len(b)
		}
		chunk := b[i:end]
		chunks = append(chunks, chunk)
	}

	return ByteMatrix(chunks)
}

type ByteMatrix [][]byte

func (b ByteMatrix) Transpose() ByteMatrix {
	r := make([][]byte, len(b[0]))
	for x, _ := range r {
		r[x] = make([]byte, len(b))
	}
	for y, s := range b {
		for x, e := range s {
			r[x][y] = e
		}
	}
	return ByteMatrix(r)
}

func Equals(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
