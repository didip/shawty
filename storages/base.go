// Package storages allows multiple implementation on how to store short URLs.
package storages

type Storage interface {
	// Hash(string) returns the short URL that would be generated from a long URL
	Code(string) string
	// Save(string) takes a full URL and returns the short URL after saving it to storage
	Save(string) (string, error)
	// Load(string) takes a short URL and returns the original full URL by retrieving it from storage
	Load(string) (string, error)
}
