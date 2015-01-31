package store

type Store interface {
    Add(key, value string) error
    Get(key string) (string, bool)
}
