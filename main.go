package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// entry point to whole app
	fmt.Println("Here we goooo")

	serve, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Conn Err TCP: ", err)
		panic("PANIC ON CONN ")
	}

	conn, err := serve.Accept()
	if err != nil {
		fmt.Println("Conn Err TCP: ", err)
		panic("PANIC ON CONN ")
	}
	for {
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("Hanldeing conn...")
	fmt.Println(conn)
	scanner := bufio.NewScanner(conn)
	fmt.Println("what about now")
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Input: ", input)
		output := fmt.Sprintf("Received: %s\n", input)
		conn.Write([]byte(output))
	}

}
