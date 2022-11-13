package golang_testing

import "sync"

type inMemoryHashtable struct {
	m   map[string][]byte
	lck sync.RWMutex
}

func NewInMemoryHashTable() Hashtable {
	return &inMemoryHashtable{m: make(map[string][]byte)}
}

func (i *inMemoryHashtable) Get(key string) ([]byte, error) {
	i.lck.RLock()
	defer i.lck.RUnlock()
	val, ok := i.m[key]
	if !ok {
		return nil, ErrNotFound
	}
	return val, nil
}

func (i *inMemoryHashtable) Set(key string, value []byte) error {
	i.lck.RLock()
	defer i.lck.RUnlock()
	i.m[key] = value
	return nil
}
