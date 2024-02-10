package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// kvStore := storage.NewKeyValueStore()

	fmt.Println("Lounching the Service...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("unimplemented %s", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("You are now connected in main fraim\n"))
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection Colsed", err.Error())
			return
		}
		if strings.TrimSpace(message) == string("exit") {
			conn.Close()
			return
		}
		fmt.Println("Message Recived: ", message, string(message), string(message) == string("exit"))
		conn.Write([]byte(message))
	}
}
