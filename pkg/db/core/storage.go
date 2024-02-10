package core

import "fmt"

type KeyValueStore struct {
	store map[string]interface{}
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]interface{}),
	}
}

func (kv *KeyValueStore) Set(key string, value interface{}) (string, interface{}) {
	kv.store[key] = value
	fmt.Println("set: ", kv.store[key])
	return key, value
}

func (kv *KeyValueStore) Get(key string) (interface{}, bool) {
	value, exists := kv.store[key]
	return value, exists
}

func (kv *KeyValueStore) Delete(key string) (string, bool) {
	fmt.Println("DEL: ", kv.store[key])
	_, ok := kv.store[key]
	if ok {
		delete(kv.store, key)
	}
	return key, ok
}
