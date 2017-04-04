package aes

import (
	"cryptopals/model"
	"cryptopals/util"
	"errors"
)

type EcbSingleByteOracle struct {
	Cipher  AesEcb
	Unknown []byte
	Key     []byte
}

func NewEcbSingleByteOracle(unknown, key []byte) EcbSingleByteOracle {
	return EcbSingleByteOracle{
		Unknown: unknown,
		Key:     key,
	}
}

func (this EcbSingleByteOracle) Encrypt(plain []byte) ([]byte, error) {
	//Now take that same function and have it append to the plaintext,
	//BEFORE ENCRYPTING, the following string:
	combined := []byte{}
	combined = append(combined, []byte(plain)...)
	combined = append(combined, this.Unknown...)

	encrypted, err := this.Cipher.Encrypt(string(combined), string(this.Key))
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

//Feed identical bytes of your-string to the function 1 at a time ---
//start with 1 byte ("A"), then "AA", then "AAA" and so on.
//Discover the block size of the cipher. You know it, but do this step anyway.
func (this EcbSingleByteOracle) FindBlockSize() (int, error) {
	a := []byte{}
	encryptedLengths := []int{}

	for i := 0; i < 2560; i++ {
		a = append(a, 'A')

		e, err := this.Encrypt(a)
		if err != nil {
			return 0, err
		}

		encryptedLengths = append(encryptedLengths, len(e))
	}

	distinct := util.Distinct(encryptedLengths)

	differences := []int{}
	for i := 0; i < len(distinct)-2; i++ {
		differences = append(differences, distinct[i+1]-distinct[i])
	}

	sizes := util.Distinct(differences)
	if len(sizes) != 1 {
		return 0, errors.New("block size not found.")
	}

	size := sizes[0]

	return size, nil
}

//Detect that the function is using ECB. You already know, but do this step anyways.
func (this EcbSingleByteOracle) IsEcb() (bool, error) {
	crib := util.GenerateCrib(2560, 'A')

	encrypted, err := this.Encrypt(crib)
	if err != nil {
		return false, err
	}

	return EcbOracle{}.IsEcb(encrypted), nil
}

func (this EcbSingleByteOracle) BuildMap(prepend []byte, blockSize, chunkNumber int) (map[int][]byte, error) {
	if len(prepend) != blockSize-1 {
		return nil, errors.New("prepend length incorrect")
	}

	charMap := map[int][]byte{}

	for i := 0; i < 256; i++ {
		appended := append(prepend, byte(i))

		encrypted, err := this.Encrypt(appended)
		if err != nil {
			return nil, err
		}

		chunks := model.ByteArray(encrypted).Chunk(blockSize)

		charMap[i] = chunks[chunkNumber]
	}

	return charMap, nil
}
