package storages

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Filesystem struct {
	Root string
	sync.RWMutex
}

func (s *Filesystem) Init(root string) error {
	s.Root = root
	return os.MkdirAll(s.Root, 0744)
}

func (s *Filesystem) Code() string {
	files, _ := ioutil.ReadDir(s.Root)
	return strconv.FormatUint(uint64(len(files)+1), 36)
}

func (s *Filesystem) Save(url string) string {
	code := s.Code()

	s.Lock()
	ioutil.WriteFile(filepath.Join(s.Root, code), []byte(url), 0744)
	s.Unlock()

	return code
}

func (s *Filesystem) Load(code string) (string, error) {
	s.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(s.Root, code))
	s.Unlock()

	return string(urlBytes), err
}
