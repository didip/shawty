// Package storages allows multiple implementation on how to store URLs as shorter names and retrieve them later.
//
// There are currently two types of storage layers, Named and Unnamed.
// Named storage layers allow the user to provide a short name and a URL
// Unnamed storage layers only accept the URL to store
package storage

import (
	"errors"
	"net/url"
)

type Storage interface {
	// Load(string) takes a short URL and returns the original full URL by retrieving it from storage
	Load(short string) (string, error)
}

type UnnamedStorage interface {
	Storage
	// Save(string) takes a full URL and returns the short URL after saving it to storage
	Save(url string) (string, error)
}

type NamedStorage interface {
	Storage
	// SaveName takes a short and a url and saves the name to use for saving a url
	SaveName(short string, url string) error
}

var (
	ErrURLEmpty   = errors.New("provided URL is of zero length")
	ErrShortEmpty = errors.New("provided short name is of zero length")

	ErrURLNotAbsolute = errors.New("provided URL is not an absolute URL")

	ErrShortNotSet = errors.New("storage layer doens't have a URL for that short code")
	ErrShortInUse  = errors.New("tried to set short, but unable to find a unique shortname within 10 tries")
)

func validateShort(short string) error {
	if short == "" {
		return ErrShortEmpty
	}

	return nil
}

func validateURL(rawURL string) (*url.URL, error) {
	if rawURL == "" {
		return nil, ErrShortEmpty
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if !parsedURL.IsAbs() {
		return nil, ErrURLNotAbsolute
	}

	return parsedURL, nil
}
