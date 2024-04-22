package base

import (
	"fmt"

	"github.com/bwlee13/gopherdb/store/request"
	"github.com/bwlee13/gopherdb/store/response"
)

const (
	LRU_TYPE  = "LRU"  // Least Recently Used
	LFU_TYPE  = "LFU"  // Least Frequently Used
	MRU_TYPE  = "MRU"  // Most Recently Used
	ARC_TYPE  = "ARC"  // Adaptive Replacement Cache
	TLRU_TYPE = "TLRU" // Time-aware Least Recently Used
)

type CommandFunc func(args request.CacheRequest) response.CacheResponse

type Store struct {
	Config   interface{}
	commands map[string]CommandFunc
}

// TODO: NewStore to read from config & take Cache Algorithm
func NewStore() *Store {
	store := &Store{
		commands: make(map[string]CommandFunc),
	}
	store.commands["get"] = store.handleGet
	store.commands["ping"] = store.handlePing
	return store
}

func (store *Store) Execute(cmd string, args request.CacheRequest) response.CacheResponse {
	if command, found := store.commands[cmd]; found {
		return command(args)
	}
	return response.CacheResponse{Error: fmt.Sprintf("Unknown command: %s", cmd)}
}

// handleGet processes the "get" command.
func (store *Store) handleGet(args request.CacheRequest) response.CacheResponse {
	fmt.Println("Command is GET: ", args)
	// Your GET command handling logic here.
	return response.CacheResponse{} // Modify according to your response structure
}

// handlePing processes the "ping" command.
func (store *Store) handlePing(args request.CacheRequest) response.CacheResponse {
	fmt.Println("Processing PING command")
	// Your PING command handling logic here.
	return response.CacheResponse{} // Modify according to your response structure
}
