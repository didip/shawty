package storage_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/shawty/storage"
)

func setupInmemStorage(t testing.TB) *storage.Inmem {
	s, err := storage.NewInmem()
	assert.Nil(t, err)

	return s
}

func TestInmemSave(t *testing.T) {
	s := setupInmemStorage(t)

	_, err := s.Save(randString(100))
	assert.Nil(t, err)
}

func TestInmemSaveName(t *testing.T) {
	s := setupInmemStorage(t)

	err := s.SaveName("shortname", "http://google.com")
	assert.Nil(t, err)
}

func TestInmemMissingLoad(t *testing.T) {
	s := setupInmemStorage(t)

	_, err := s.Load("non-existant-short-string")
	assert.NotNil(t, err)

	short, err := s.Save("http://google.com")
	assert.Nil(t, err)

	_, err = s.Load(short)
	assert.Nil(t, err)

	_, err = s.Load(short + short)
	assert.NotNil(t, err)
	assert.Equal(t, err, storage.ErrCodeNotSet)
}

func TestInmemLoad(t *testing.T) {
	s := setupInmemStorage(t)

	url := randString(100)

	short, err := s.Save(url)
	assert.Nil(t, err)

	long, err := s.Load(short)
	assert.Nil(t, err)

	assert.Equal(t, url, long)
}

func TestInmemMultipleLoads(t *testing.T) {
	s := setupInmemStorage(t)

	urls := make([]string, 10)
	shorts := make([]string, 10)
	var err error
	for i := 0; i < len(urls); i++ {
		urls[i] = randString(100)
		shorts[i], err = s.Save(urls[i])
		t.Logf("Saved '%s' to '%s'", urls[i], shorts[i])
		assert.Nil(t, err)
	}

	for range urls {
		i := rand.Intn(len(urls))
		t.Logf("Checking %v: %s", i, shorts[i])
		long, err := s.Load(shorts[i])
		assert.Nil(t, err)

		assert.Equal(t, urls[i], long)
	}
}
