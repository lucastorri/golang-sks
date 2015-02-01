package store

type MemStore struct {
    values map[string]string
}

func NewMemStore() *MemStore {
    return &MemStore{ make(map[string]string) }
}

func (m *MemStore) Add(key, value string) error {
    m.values[key] = value
    return nil
}

func (m *MemStore) Get(key string) (string, bool) {
    value, ok := m.values[key]
    return value, ok
}
