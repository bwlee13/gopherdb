package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwlee13/gopherdb/storage/base"
	"github.com/pkg/errors"
)

var (
	store *base.Store
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

func (service *Service) Start() error {
	if err := StartServer(service); err != nil {
		return err
	}
	return nil
}

func StartServer(service *Service) (err error) {
	serve, err := net.Listen("tcp", service.addr)
	if err != nil {
		return errors.Wrap(err, "tcp listen err")
	}

	defer serve.Close()

	for {
		conn, err := serve.Accept()
		if err != nil {
			log.Println("Conn Err TCP: ", err)
			continue
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
			log.Printf("Read ERR: %s \n", err)
			return
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

func waitForShutdown(shutdownChan chan struct{}) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigChan:
		log.Println("User is shutting down server...")
	case <-shutdownChan:
		log.Println("Internal error, shutting down server...")
	}
}

func InitServer(port string, policy string) {
	log.Println("Port is set to: ", port)
	log.Println("Policy is set to: ", policy)

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	store = base.NewStore(policy)
	service := NewService(port, store)
	shutdownChan := make(chan struct{})

	log.Println("Starting server...")
	go func() {
		if err := service.Start(); err != nil {
			log.Println("Error: ", err)
			shutdownChan <- struct{}{}
		}
	}()

	waitForShutdown(shutdownChan)
}
