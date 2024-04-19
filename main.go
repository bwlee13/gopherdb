package main

import (
	"fmt"
	"gopherdb"
	"net"
)

func main() {
	// entry point to whole app
	fmt.Println("Here we goooo")
	gopherdb.HandleConn()
	// testing

	serve, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Conn Err TCP: ", err)
		panic("PANIC ON CONN ")
	}

	for {
		conn, err := serve.Accept()

		if err != nil {
			fmt.Println("Conn Err TCP: ", err)
			panic("PANIC ON CONN ")
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("Hanldeing conn...")
	fmt.Println(conn)
}
