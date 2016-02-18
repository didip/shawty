package storage

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var randCharsLen = len(randChars)

func getRandomString(length int) string {
	s := make([]rune, length)

	for i := range s {
		s[i] = randChars[rand.Intn(randCharsLen)]
	}

	return string(s)
}
