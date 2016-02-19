// Package storages allows multiple implementation on how to store short URLs.
package storage

import "errors"

type Storage interface {
	// Load(string) takes a short URL and returns the original full URL by retrieving it from storage
	Load(string) (string, error)
}

type UnnamedStorage interface {
	Storage
	// Save(string) takes a full URL and returns the short URL after saving it to storage
	Save(url string) (string, error)
}

type NamedStorage interface {
	Storage
	// SaveName takes a name and a url and saves the name to use for saving a url
	SaveName(name string, url string) error
}

var (
	ErrCodeInUse  = errors.New("tried to set short, but unable to find a unique shortname")
	ErrCodeNotSet = errors.New("storage layer doens't have a url for that short code")
	ErrURLEmpty   = errors.New("provided url is empty")
	ErrNameEmpty  = errors.New("provided short name is empty")
)
