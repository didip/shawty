package storage_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thomaso-mirodin/shawty/storage"
)

func setupFilesystemStorage(t testing.TB) storage.UnnamedStorage {
	dir, err := ioutil.TempDir("", "BenchmarkFilesystem")
	require.Nil(t, err)

	s, err := storage.NewFilesystem(dir)
	require.Nil(t, err)

	return s
}
