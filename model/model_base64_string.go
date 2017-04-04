package model

type Base64String string

func NewBase64String(h string) Base64String {
	return Base64String(h)
}
