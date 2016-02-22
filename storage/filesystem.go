package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Filesystem struct {
	Root string
	c    uint64
	mu   sync.RWMutex
}

func NewFilesystem(root string) (*Filesystem, error) {
	s := &Filesystem{
		Root: root,
	}
	return s, os.MkdirAll(s.Root, 0744)
}

func (s *Filesystem) Code(url string) string {
	return strconv.FormatUint(s.c, 36)
}

func (s *Filesystem) Save(url string) (string, error) {
	if url == "" {
		return "", ErrURLEmpty
	}

	code := s.Code(url)

	s.mu.Lock()
	err := ioutil.WriteFile(filepath.Join(s.Root, code), []byte(url), 0744)
	if err == nil {
		s.c++
	}
	s.mu.Unlock()

	return code, err
}

func (s *Filesystem) Load(code string) (string, error) {
	if code == "" {
		return "", ErrNameEmpty
	}

	s.mu.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(s.Root, code))
	s.mu.Unlock()

	return string(urlBytes), err
}
