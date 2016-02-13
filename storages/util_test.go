package storages_test

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/shawty/storages"
)

func setupFilesystemStorage(t testing.TB) storages.Storage {
	dir, err := ioutil.TempDir("", "BenchmarkFilesystem")
	assert.Nil(t, err)

	s, err := storages.NewFilesystem(dir)
	assert.Nil(t, err)

	return s
}

func randString(length int) string {
	b := make([]byte, length)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const maxLetters = len(letters)

	for i := range b {
		b[i] = letters[rand.Intn(maxLetters)]
	}

	return string(b)
}
