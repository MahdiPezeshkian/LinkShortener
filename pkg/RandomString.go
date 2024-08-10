package pkg

import (
	"time"

	"golang.org/x/exp/rand"
)

func RandomString(minLen, maxLen int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)

	length := r.Intn(maxLen-minLen+1) + minLen

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}