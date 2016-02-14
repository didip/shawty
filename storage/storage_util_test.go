package storage_test

import "math/rand"

func randString(length int) string {
	b := make([]byte, length)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const maxLetters = len(letters)

	for i := range b {
		b[i] = letters[rand.Intn(maxLetters)]
	}

	return string(b)
}
