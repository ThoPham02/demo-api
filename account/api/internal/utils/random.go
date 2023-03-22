package utils

import (
	"math/rand"
	"time"
)

var (
	alpha = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomNumber(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RandomString(n int) string {
	result := []byte{}
	alphaLen := len(alpha)
	for i := 0; i < n; i++ {
		byte := alpha[RandomNumber(0, int64(alphaLen))]
		result = append(result, byte)
	}
	return string(result)
}
