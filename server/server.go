package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwlee13/gopherdb/storage/base"
	"github.com/pkg/errors"
)

var (
	listen = flag.String("listen", ":42069", "address to listen to")
	store  *base.Store
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

func waitForShutdown() {
	t := make(chan os.Signal, 1)
	signal.Notify(t, os.Interrupt, syscall.SIGTERM)
	<-t

	log.Println("Shutting down server...")
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

func InitServer() {
	flag.Parse()

	// default LRU because thats all i have right now
	store = base.NewStore("LRU")
	service := NewService(*listen, store)

	go service.Start()
	log.Println("Starting server...")
	waitForShutdown()

}
