package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	// entry point to whole app
	fmt.Println("Here we goooo")

	serve, _ := net.Listen("tcp", ":42069")

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
			fmt.Println("Error reading length: ", err)
			return
		}
		fmt.Println("Length? ", length)

		payload := make([]byte, 250)
		fmt.Println("Payload? ", payload)
		if _, err := conn.Read(payload); err != nil {
			fmt.Println("Error reading payload: ", err)
			return
		}

		fmt.Println("Received payload: ", payload)

		if _, err := conn.Write([]byte("Ack\n")); err != nil {
			fmt.Println("Error writing to connection: ", err)
			return
		}
	}

}
