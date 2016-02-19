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

func (s *Filesystem) Init(root string) error {
	s.Root = root
	return os.MkdirAll(s.Root, 0744)
}

func NewFilesystem(root string) (*Filesystem, error) {
	s := new(Filesystem)
	return s, s.Init(root)
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
