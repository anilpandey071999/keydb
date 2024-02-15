package main

import (
	"fmt"

	"github.com/anilpandey071999/keydb/pkg/sdk/client"
)

func main() {
	keyDB := client.NewClient("0.0.0.0", "8080")
	fmt.Println(keyDB)
	data, err := keyDB.SET("hello", 125)
	fmt.Println("SET: ", data, err)
	data, err = keyDB.GET("hello")
	fmt.Println("GET: ", data)
	data, err = keyDB.GET("hello")
	fmt.Println("GET: ", data, err)
	data, err = keyDB.DELETE("hello")
	fmt.Println("DELETE: ", data, err)
	data, err = keyDB.GET("hello")
	fmt.Println("GET: ", data, err)
	keyDB.Exit()
}
