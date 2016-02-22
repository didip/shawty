package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// CleanPath removes any path transversal nonsense
func CleanPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return filepath.Clean(path)
}

// Takes a possibly multilevel path and flattens it by dropping any slashes
func FlattenPath(path string, separator string) string {
	return strings.Replace(path, string(os.PathSeparator), separator, -1)
}

func (s *Filesystem) SaveName(short, long string) error {
	if short == "" {
		return ErrNameEmpty
	}
	if long == "" {
		return ErrURLEmpty
	}

	short = FlattenPath(CleanPath(short), "_")

	s.mu.Lock()
	err := ioutil.WriteFile(filepath.Join(s.Root, short), []byte(long), 0744)
	if err == nil {
		s.c++
	}
	s.mu.Unlock()

	return err
}

func (s *Filesystem) Load(code string) (string, error) {
	if code == "" {
		return "", ErrNameEmpty
	}

	code = FlattenPath(CleanPath(code), "_")

	s.mu.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(s.Root, code))
	s.mu.Unlock()

	if _, ok := err.(*os.PathError); ok {
		return "", ErrCodeNotSet
	}

	return string(urlBytes), err
}
