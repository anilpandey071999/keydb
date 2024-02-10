package storage

type KeyValueStore struct {
	store map[string]interface{}
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]interface{}),
	}
}

func (kv *KeyValueStore) Set(key string, value interface{}) {
	kv.store[key] = value
}

func (kv *KeyValueStore) Get(key string) (interface{}, bool) {
	value, exists := kv.store[key]
	return value, exists
}

func (kv *KeyValueStore) Delete(key string) {
	delete(kv.store, key)
}
