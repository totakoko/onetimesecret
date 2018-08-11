package helpers

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	// initialize the random generator
	rand.Seed(time.Now().UTC().UnixNano())
}

func GenerateRandomString(n int) string {
	if n <= 0 {
		panic("Rand(): i must be > 0")
	}
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(bytes)
}
