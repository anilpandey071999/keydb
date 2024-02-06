package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// kvStore := storage.NewKeyValueStore()

	fmt.Println("Lounching the Service...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("unimplemented %s", err.Error())
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection Colsed", err.Error())
			return
		}
		fmt.Println("Message Recived: ", message, string(message))
		conn.Write([]byte(message))
	}
}
