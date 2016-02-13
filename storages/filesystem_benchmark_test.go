package storages_test

import (
	"math"
	"testing"
)

func BenchmarkFilesystemEmptyCode(b *testing.B) {
	s := setupFilesystemStorage(b)

	url := randString(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Code(url)
	}
}

func BenchmarkFilesystemSmallCode(b *testing.B) {
	s := setupFilesystemStorage(b)

	urls := make([]string, 100)
	urlCount := len(urls)

	for i := 0; i < urlCount; i++ {
		urls[i] = randString(100)
		s.Save(urls[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Code(urls[i%urlCount])
	}
}

func BenchmarkFilesystemLotsCode(b *testing.B) {
	s := setupFilesystemStorage(b)

	urls := make([]string, int(math.Pow(10, 4)))
	urlCount := len(urls)

	for i := 0; i < urlCount; i++ {
		urls[i] = randString(100)
		s.Save(urls[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Code(urls[i%urlCount])
	}
}
