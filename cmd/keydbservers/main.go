package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/anilpandey071999/keydb/pkg/db/core"
	"github.com/anilpandey071999/keydb/pkg/db/persistence"
)

func main() {
	fmt.Println("Lounching the Service...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("unimplemented: ", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// conn.Write([]byte("You are now connected in main fraim"))
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection Colsed", err.Error())
			return
		}
		if message == "" {
			conn.Write([]byte(err.Error()))
		}
		if strings.TrimSpace(message) == string("exit") {
			conn.Close()
			return
		}
		result, err := core.QueryProcesser(message)
		if err != nil {
			conn.Write([]byte(err.Error()))
		}

		persistence.WriteAOF(message)

		jsonData, err := json.Marshal(result)
		if err != nil {
			conn.Write([]byte(err.Error()))
		}
		// return []string{string(jsonData)}, nil
		fmt.Println("Message Recived: ", message, string(message), string(message) == string("exit"))
		conn.Write([]byte(jsonData))
	}
}
