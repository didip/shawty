package storage

import "sync"

type Inmem struct {
	RandLength int

	m  map[string]string
	mu sync.RWMutex
}

func NewInmem() (*Inmem, error) {
	s := &Inmem{
		RandLength: 8,

		m: make(map[string]string),
	}
	return s, nil
}

func (s *Inmem) Save(url string) (string, error) {
	if url == "" {
		return "", ErrURLEmpty
	}

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
	if code == "" {
		return ErrNameEmpty
	}
	if url == "" {
		return ErrURLEmpty
	}

	s.mu.Lock()
	s.m[code] = url
	s.mu.Unlock()
	return nil
}

func (s *Inmem) Load(code string) (string, error) {
	if code == "" {
		return "", ErrNameEmpty
	}

	s.mu.Lock()
	url, ok := s.m[code]
	s.mu.Unlock()
	if !ok {
		return "", ErrCodeNotSet
	}

	return url, nil
}
