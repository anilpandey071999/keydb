package client

import (
	"encoding/json"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(address, port string) *Client {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		panic(err)
	}

	return &Client{conn: conn}
}

func (c *Client) SET(key string, value interface{}) (interface{}, error) {
	jsonData, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	c.conn.Write([]byte("SET " + key + " " + string(jsonData) + "\n"))
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	response := string(buffer[:n])
	return response, nil
}

func (c *Client) GET(key string) (interface{}, error) {
	_, err := c.conn.Write([]byte("GET " + key + "\n"))
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	response := string(buffer[:n])
	return response, nil
}

func (c *Client) DELETE(key string) (interface{}, error) {
	_, err := c.conn.Write([]byte("DELETE " + key + "\n"))
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	response := string(buffer[:n])
	return response, nil
}

func (c *Client) Exit() {
	_, err := c.conn.Write([]byte("exit"))
	if err != nil {
		panic(err)
	}
}
