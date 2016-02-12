package storages_test

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/thomaso-mirodin/shawty/storages"
)

func BenchmarkCode(b *testing.B) {
	dir, _ := homedir.Dir()
	storage := &Filesystem{}
	storage.Init(filepath.Join(dir, "shawty"))

	for i := 0; i < b.N; i++ {
		storage.Code()
	}
}
