//doesn't necessarily need to be in aes
package aes

type BlockCipherMode string

const (
	ECB BlockCipherMode = "ECB" //electronic code book
	CBC                 = "CBC" //cipher block chaining
)
