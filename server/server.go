package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bwlee13/gopherdb/storage/base"
	"github.com/pkg/errors"
)

type Service struct {
	addr  string
	store *base.Store
}

func NewService(addr string, store *base.Store) *Service {
	return &Service{
		addr:  addr,
		store: store,
	}
}

func (service *Service) Start() {
	StartServer(service)
}

func StartServer(service *Service) (err error) {
	serve, err := net.Listen("tcp", service.addr)
	if err != nil {
		fmt.Printf("Failed to bind to port: %s", service.addr)
		return errors.Wrap(err, "listen")
	}

	defer serve.Close()

	for {
		conn, err := serve.Accept()
		if err != nil {
			fmt.Println("Conn Err TCP: ", err)
			panic("PANIC ON CONN")
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		// var length uint32
		// if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
		// 	if err != io.EOF {
		// 		fmt.Println("Error reading length:", err)
		// 	}
		// 	break
		// }

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read ERR: %s \n", err)
			panic("Read Err")
		}

		fmt.Println("Length:", len(buf[:n]))

		log.Printf("command:\n %s", buf[:n])
		fmt.Println(len(buf[:n]))
		payload := buf[:n]
		fmt.Println("Received payload:", payload)

		// msg, _ := Parse(string(buf[:n]))
		// if msg == nil {
		// 	fmt.Println("Error when parsing message")
		// 	return errors.Wrap(err, "parse command")
		// }

		if _, err := conn.Write([]byte("Ack\n")); err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}

}
