package helpers

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Rand(n int) string {
	if n <= 0 {
		panic("Rand(): i must be > 0")
	}
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(bytes)
}
