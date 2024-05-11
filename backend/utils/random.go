package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const digits = "0123456789"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomAlphabetString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUserName() string {
	return RandomAlphabetString(6)
}

func RandomPrice() int64 {
	return RandomInt(0, 1000)
}

func RandomDiscount() int64 {
	return RandomInt(0, 100)
}

func RandomLanguage() string {
	currencies := []string{"chn", "jp"}
	return currencies[rand.Intn(len(currencies))]
}

func RandomNumberString(n int) string {
	var sb strings.Builder
	k := len(digits)

	for i := 0; i < n; i++ {
		c := digits[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
