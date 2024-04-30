package base

import (
	"fmt"

	"github.com/bwlee13/gopherdb/storage/cache"
	"github.com/bwlee13/gopherdb/storage/lru"
	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/bwlee13/gopherdb/storage/response"
)

const (
	STORE_GET         = "get"
	STORE_PING        = "ping"
	STORE_PUT         = "put"
	STORE_ADD         = "add"
	STORE_DELETE      = "delete"
	STORE_FLUSH       = "flush"
	STORE_NODE_SIZE   = "nodeSize"
	STORE_APP_METRICS = "getAppMetrics"
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
	policy   string
	Cache    cache.Cache
	commands map[string]interface{}
}

// TODO: NewStore to read from config & take Cache Algorithm
func NewStore(policy string) *Store {
	store := &Store{
		policy:   policy,
		commands: make(map[string]interface{}),
	}
	// store.commands["get"] = store.handleGet
	// store.commands["ping"] = store.handlePing
	return store
}

func (store *Store) BuildStore() {
	store.Cache = store.newCacheFromPolicy(store.policy)
	store.commands = store.registerHandlers()
}

func (store *Store) newCacheFromPolicy(policy string) cache.Cache {
	switch policy {
	case LRU_TYPE:
		return lru.NewLRU()
	default:
		return nil
	}
}

func (baseStore *Store) registerHandlers() map[string]interface{} {
	return map[string]interface{}{
		STORE_GET:  baseStore.Cache.Get,
		STORE_PUT:  baseStore.Cache.Put,
		STORE_PING: baseStore.Cache.Ping,
	}

}

func (store *Store) Execute(cmd string, args request.CacheRequest) response.CacheResponse {
	fmt.Println("cmd rec: ", cmd)
	fmt.Println("what is policy", store.policy)
	fmt.Println("store cache ", store.Cache)
	fmt.Println("store commands ", store.commands)
	if _, found := store.commands[cmd]; found {
		// Convert the interface{} type to the specific function type using a variable.
		// some weird Go shit I haven't seen before. Kinda cool tho
		// get func from map, assert into func type, call with args
		res := store.commands[cmd].(func(request.CacheRequest) response.CacheResponse)(args)
		return res
	}
	fmt.Println("why not found...")
	fmt.Println("store: ", store.commands[cmd])
	fmt.Println("store ping: ", store.commands["ping"])
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
