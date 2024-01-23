package helper

import (
	"math/rand"
	"time"
)

type CharsetType int8

const (
	CharsetAlphaNumeric CharsetType = iota
	CharsetAlpha
	CharsetNumeric
)

const charsetAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumeric = "0123456789"

var Charsets = map[CharsetType]string{
	CharsetAlphaNumeric: charsetAlpha + charsetNumeric,
	CharsetAlpha:        charsetAlpha,
	CharsetNumeric:      charsetNumeric,
}

func GenerateRandomString(charsetType CharsetType, length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := Charsets[charsetType]

	b := make([]byte, GetDefaultNumber[int](length, 10))
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomAplphaNumeric(length int) string {
	return GenerateRandomString(CharsetAlphaNumeric, length)
}

func RandomAplpha(length int) string {
	return GenerateRandomString(CharsetAlpha, length)
}

func RandomNumeric(length int) string {
	return GenerateRandomString(CharsetNumeric, length)
}
