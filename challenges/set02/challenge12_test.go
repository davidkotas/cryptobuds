package set02

import (
	"cryptopals/aes"
	"cryptopals/model"
	"cryptopals/util"
	"encoding/base64"
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_Byte_at_a_time_ECB_decryption_Simple(t *testing.T) {
	//assign a single random key, once, to a global variable
	key := []byte{127, 223, 59, 114, 11, 201, 90, 31, 221, 164, 71, 1, 136, 229, 194, 82}

	unknown, err := readUnknown()
	if err != nil {
		t.Fatal(err)
	}

	oracle := aes.NewEcbSingleByteOracle(unknown, key)

	isEcb, err := oracle.IsEcb()
	if err != nil {
		t.Fatal(err)
	}

	if !isEcb {
		t.Fatal("block mode not ecb")
	}

	blockSize, err := oracle.FindBlockSize()
	if err != nil {
		t.Fatal(err)
	}

	if blockSize != 16 {
		t.Fatal("block size not 16")
	}

	unknownEncrypted, err := oracle.Encrypt([]byte{})
	if err != nil {
		t.Fatal(err)
	}

	log.Println(len(unknownEncrypted))

	unknownChunks := model.ByteArray(unknownEncrypted).Chunk(blockSize)

	log.Println(len(unknownChunks))

	truncated := util.GenerateCrib(blockSize-1, 'A')

	//Knowing the block size, craft an input block that is exactly 1 byte short
	//(for instance, if the block size is 8 bytes, make "AAAAAAA").
	//Think about what the oracle function is going to put in that last byte position.
	// --> the oracle is going to put the first byte of the unknown string
	output, err := oracle.Encrypt(truncated)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(output[0:16])

	charMap, err := oracle.BuildMap(truncated, blockSize, 0)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range charMap {
		if model.Equals(output[0:16], v) {
			log.Printf("MATCH %v", k)
			next := v[1:16]
			log.Println(next)
		}
	}

	log.Println("challenge 12 - not implemented!")
}

func readUnknown() ([]byte, error) {
	base64Bytes, err := ioutil.ReadFile("challenge12_text.txt")
	if err != nil {
		return nil, err
	}

	unknown, err := base64.StdEncoding.DecodeString(string(base64Bytes))
	if err != nil {
		return nil, err
	}

	return unknown, nil
}
