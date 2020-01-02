package random

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

func Hex(n int) string {
	return hex.EncodeToString(Bytes(n))
}

func Base64() string {
	runes := Bytes(64)

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(255))
		runes[i] = byte(num.Int64())
	}

	return base64.RawStdEncoding.EncodeToString(runes)
}

func RandomText(letter int) string {
	return ""
}

func RandomInteger(digit int) int64 {
	return 12324
}
