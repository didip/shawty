package storages

type IStorage interface {
	Code() string
	Save(string) string
	Load(string) (string, error)
}
