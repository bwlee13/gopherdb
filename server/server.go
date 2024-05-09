package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwlee13/gopherdb/storage/base"
	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/bwlee13/gopherdb/storage/response"
	"github.com/pkg/errors"
)

var (
	store *base.Store
)

type Service struct {
	addr  string
	store *base.Store
	quit  chan bool
}

func NewService(addr string, store *base.Store) *Service {
	return &Service{
		addr:  addr,
		store: store,
		quit:  make(chan bool),
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
	go func() error {
		for {
			conn, err := serve.Accept()
			if err != nil {
				select {
				case <-service.quit:
					log.Println("Server requested to shut down...")
					return fmt.Errorf("shutdown req")
				default:
					log.Println("Connection accept error: ", err)
					return fmt.Errorf("shutdown req")
				}
			}

			defer conn.Close()
			go handleConn(conn, service)
		}
	}()

	<-service.quit

	return fmt.Errorf("Quitting")
}

func handleConn(conn net.Conn, service *Service) error {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client has closed connection")
				return fmt.Errorf("client closed conn")
			} else {
				log.Printf("Read ERR: %s \n", err)
				return fmt.Errorf("read Err: %s", err)
			}
		}

		defer conn.Close()

		payload := buf[:n]

		message := strings.TrimSpace(string(payload))
		fmt.Println("Received message:", message)
		// @TODO Pass message to Parser() func
		// @TODO Pass Parsed MSG to HandleCmd() func
		// @TODO Return response to client fro HanldeCmd() -> {cmdName}Cmd()

		// obv find a way better way than this. Maybe use handler and pass trimmed message
		// also do switch case, not if elif...
		if message == "ping" {
			fmt.Println("PING REC")
			res := HandleCommands("ping")
			fmt.Println("RESULT: ", res)
			msg := res.Message + "\n"
			fmt.Println("sending message: ", msg)
			encoder := json.NewEncoder(conn)
			err = encoder.Encode(res)
			if err != nil {
				fmt.Println("err: ", err)
			}

			// if _, err := conn.Write(res); err != nil {
			// 	fmt.Println("Error writing to connection:", err)
			// 	return fmt.Errorf("error writing to connection: %s", err)
			// }
		} else if message == "shutdown" {
			fmt.Println("Shutdown command received, closing connection.")
			service.quit <- true
			return fmt.Errorf("shutdown")
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
	store.BuildStore()
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

func HandleCommands(command string) response.CacheResponse {
	if command == "ping" {
		return store.Execute("ping", request.NewEmptyCacheRequest())
	}
	return response.NewPingResponse()
}

func StopServer(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	conn, err := net.Dial("tcp", port)
	if err != nil {
		return errors.Wrap(err, "Stop Conn Fail")
	}
	defer conn.Close()

	_, err = conn.Write([]byte("shutdown"))
	if err != nil {
		return errors.Wrap(err, "Shutdown Fail")
	}
	return nil
}
