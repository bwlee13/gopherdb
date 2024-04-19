package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:42069")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	sendString(conn, "hello world")
}

func sendString(conn net.Conn, s string) {
	// Convert the string to bytes
	data := []byte(s)

	// Send the length of the data as a uint32
	err := binary.Write(conn, binary.BigEndian, uint32(len(data)))
	if err != nil {
		fmt.Println("Error writing length:", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return
	}

	ack := make([]byte, 4)
	if _, err = conn.Read(ack); err != nil {
		fmt.Println("Error reading acknowledgment:", err)
		return
	}

	fmt.Println("Acknowledgment received:", string(ack))
}
