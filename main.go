package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	StartServer()
}

func StartServer() {
	// entry point to whole app
	fmt.Println("Here we goooo")

	serve, _ := net.Listen("tcp", ":42069")
	defer serve.Close()

	for {
		conn, err := serve.Accept()
		if err != nil {
			fmt.Println("Conn Err TCP: ", err)
			panic("PANIC ON CONN ")
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println(conn)
	defer conn.Close()

	for {
		var length uint32
		if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
			if err != io.EOF {
				fmt.Println("Error reading length:", err)
			}
			break
		}

		fmt.Println("Length:", length)

		if length > 1024*1024 {
			fmt.Println("Payload length too large:", length)
			return
		}
		payload := make([]byte, length)

		_, err := io.ReadFull(conn, payload)
		if err != nil {
			fmt.Println("Error reading payload:", err)
			return
		}

		fmt.Println("Received payload:", payload)

		if _, err := conn.Write([]byte("Ack\n")); err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}

}
