package storage

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var randChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var randCharsLen = len(randChars)

func getRandomString(length int) string {
	s := make([]rune, length)

	for i := range s {
		s[i] = randChars[rand.Intn(randCharsLen)]
	}

	return string(s)
}

type Inmem struct {
	RandLength int

	m  map[string]string
	mu sync.RWMutex
}

func (s *Inmem) Init() error {
	s.mu.Lock()
	s.m = make(map[string]string)
	s.RandLength = 8
	s.mu.Unlock()
	return nil
}

func NewInmem() (*Inmem, error) {
	s := new(Inmem)
	return s, s.Init()
}

func (s *Inmem) Save(url string) (string, error) {
	var code string

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 0; i < 10; i++ {
		code = getRandomString(8)

		if _, ok := s.m[code]; !ok {
			s.m[code] = url
			return code, nil
		}
	}

	return "", ErrCodeInUse
}

func (s *Inmem) SaveName(code string, url string) error {
	s.mu.Lock()
	s.m[code] = url
	s.mu.Unlock()
	return nil
}

func (s *Inmem) Load(code string) (string, error) {
	s.mu.Lock()
	url, ok := s.m[code]
	s.mu.Unlock()
	if !ok {
		return "", ErrCodeNotSet
	}

	return url, nil
}
