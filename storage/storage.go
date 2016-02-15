// Package storages allows multiple implementation on how to store short URLs.
package storage

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
