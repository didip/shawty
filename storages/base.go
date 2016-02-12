// Package storages allows multiple implementation on how to store short URLs.
package storages

type Storage interface {
	Code() string
	Save(string) string
	Load(string) (string, error)
}
