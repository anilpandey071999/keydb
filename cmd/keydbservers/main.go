package main

import (
	"fmt"

	"github.com/anilpandey071999/keydb/pkg/storage"
)

func main() {
	kvStore := storage.NewKeyValueStore()

	var key string
	var values string

	fmt.Println("you Key")
	fmt.Scanln(&key)
	
	fmt.Println("you Value")
	fmt.Scanln(&values)
	
	kvStore.Set(key, values)
	
	value, exists := kvStore.Get(key)
	if exists {
		fmt.Println("Retrived:", value)
	}
	
	kvStore.Delete(key)
	
	_, exists = kvStore.Get(key)
	if !exists {
		fmt.Println("Key1 Deleted")
	}

}
