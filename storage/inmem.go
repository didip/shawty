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
	if _, err := validateURL(url); err != nil {
		return "", err
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

	return "", ErrShortInUse
}

func (s *Inmem) SaveName(short string, url string) error {
	if err := validateShort(short); err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	s.mu.Lock()
	s.m[short] = url
	s.mu.Unlock()
	return nil
}

func (s *Inmem) Load(short string) (string, error) {
	if err := validateShort(short); err != nil {
		return "", err
	}

	s.mu.Lock()
	url, ok := s.m[short]
	s.mu.Unlock()
	if !ok {
		return "", ErrShortNotSet
	}

	return url, nil
}
