package storage_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/shawty/storage"
)

func randString(length int) string {
	b := make([]byte, length)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const maxLetters = len(letters)

	for i := range b {
		b[i] = letters[rand.Intn(maxLetters)]
	}

	return string(b)
}

func saveSomething(s storage.Storage) (short string, long string, err error) {
	unnamed, unnamedOk := s.(storage.UnnamedStorage)
	named, namedOk := s.(storage.NamedStorage)

	short = randString(10)
	long = randString(20)

	if unnamedOk {

		short, err := unnamed.Save(long)
		return short, long, err
	} else if namedOk {
		err := named.SaveName(short, long)
		return short, long, err
	} else {
		return "", "", fmt.Errorf("Storage isn't named or unnamed, can't save anything")
	}
}

var storageSetups = map[string]func(testing.TB) storage.Storage{
	"Inmem": setupInmemStorage,
	"S3":    setupS3Storage,
}

var storageCleanup = map[string]func() error{
	"S3": cleanupS3Storage,
}

func TestMain(m *testing.M) {
	res := m.Run()

	for _, cf := range storageCleanup {
		err := cf()
		if err != nil {
			log.Println("Cleanup error:", err)
		}
	}

	os.Exit(res)
}

func TestUnnamedStorageSave(t *testing.T) {
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		unnamedStorage, ok := setupStorage(t).(storage.UnnamedStorage)

		if assert.True(t, ok, name) {
			code, err := unnamedStorage.Save(testURL)
			t.Logf("[%s] unnamedStorage.Save(\"%s\") -> %#v", name, testURL, code)
			assert.Nil(t, err, name)
		}
	}
}

func TestNamedStorageSave(t *testing.T) {
	testCode := "test-named-url"
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		namedStorage, ok := setupStorage(t).(storage.NamedStorage)

		if assert.True(t, ok, name) {
			err := namedStorage.SaveName(testCode, testURL)
			t.Logf("[%s] namedStorage.SaveName(\"%s\", \"%s\") -> %#v", name, testCode, testURL, err)
			assert.Nil(t, err, name)
		}
	}
}

func TestMissingLoad(t *testing.T) {
	testCode := "non-existant-short-string"

	for name, setupStorage := range storageSetups {
		long, err := setupStorage(t).Load(testCode)
		t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, testCode, long, err)
		assert.NotNil(t, err, name)
		assert.Equal(t, err, storage.ErrCodeNotSet, name)
	}
}

func TestLoad(t *testing.T) {
	for name, setupStorage := range storageSetups {
		s := setupStorage(t)

		short, long, err := saveSomething(s)
		t.Logf("[%s] saveSomething(s) -> %#v, %#v, %#v", name, short, long, err)
		assert.Nil(t, err, name)

		newLong, err := s.Load(short)
		t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, short, long, err)
		assert.Nil(t, err, name)

		assert.Equal(t, long, newLong, name)
	}
}

// func TestNamedStorageNames(t *testing.T) {
// 	var shortNames map[string]error = map[string]error{
// 		"simple": nil,
// 		""
// 	}

// 	for name, setupStorage := range storageSetups {
// 		namedStorage, ok := setupStorage(t).(storage.NamedStorage)

// 		namedStorage.SaveName(name, url)
// 	}
// }
