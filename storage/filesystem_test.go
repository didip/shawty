package storages_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/shawty/storages"
)

func setupFilesystemStorage(t testing.TB) storages.UnnamedStorage {
	dir, err := ioutil.TempDir("", "BenchmarkFilesystem")
	assert.Nil(t, err)

	s, err := storages.NewFilesystem(dir)
	assert.Nil(t, err)

	return s
}

func TestFilesystemSave(t *testing.T) {
	s := setupFilesystemStorage(t)

	_, err := s.Save(randString(100))
	assert.Nil(t, err)
}

func TestFilesystemLoad(t *testing.T) {
	s := setupFilesystemStorage(t)

	url := randString(100)

	short, err := s.Save(url)
	assert.Nil(t, err)

	long, err := s.Load(short)
	assert.Nil(t, err)

	assert.Equal(t, url, long)
}

func TestFilesystemCode(t *testing.T) {

	s := setupFilesystemStorage(t)

	assert.Equal(t, s.Code(randString(10)), "0")
	assert.Equal(t, s.Code(randString(10)), "0")

	s.Save(randString(10))

	assert.Equal(t, s.Code(randString(10)), "1")
	assert.Equal(t, s.Code(randString(10)), "1")

	for i := 0; i < 1000; i++ {
		s.Save(randString(10))
	}

	assert.Equal(t, s.Code(randString(10)), "rt")

}

func TestFilesystemMultipleLoads(t *testing.T) {
	s := setupFilesystemStorage(t)
	fmt.Println(s.(*storages.Filesystem).Root)

	urls := make([]string, 100)
	shorts := make([]string, 100)
	var err error
	for i := 0; i < len(urls); i++ {
		urls[i] = randString(10)
		t.Logf("Saving '%s' to '%s'", urls[i], s.Code(urls[i]))
		shorts[i], err = s.Save(urls[i])
		t.Logf("Saved'%s' to '%s'", urls[i], shorts[i])
		assert.Nil(t, err)
	}

	for range urls {
		i := rand.Intn(len(urls))
		t.Log("Checking i:", i)
		long, err := s.Load(shorts[i])
		assert.Nil(t, err)

		assert.Equal(t, urls[i], long)
	}
}
