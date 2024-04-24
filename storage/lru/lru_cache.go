package lru

import (
	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/bwlee13/gopherdb/storage/response"
)

const (
	CACHE_MISS = "CACHE_MISS"
	STORED     = "STORED"
)

type LRUCache struct {
	Size      int32
	Count     int32
	Hashtable map[string]interface{}
}

func NewLRU() *LRUCache {
	return &LRUCache{
		Size:      10,
		Count:     int32(0),
		Hashtable: newHashtable(),
	}
}

func newHashtable() map[string]interface{} {
	return make(map[string]interface{})
}

func (cache *LRUCache) Get(args request.CacheRequest) response.CacheResponse {
	key := args.CacheObj.Key
	node := cache.Hashtable[key]
	if node == nil {
		return response.NewCacheMissResponse()
	}
	return response.NewResponseFromValue(node)
}

func (cache *LRUCache) Put(args request.CacheRequest) response.CacheResponse {
	key := args.CacheObj.Key
	val := args.CacheObj.Value

	cache.Hashtable[key] = val
	return response.NewResponseFromMessage(STORED, 1)
}

func (cache *LRUCache) Add(args request.CacheRequest) response.CacheResponse {
	return response.NewResponseFromMessage("METHODNOTAVAIL", 0)
}

func (cache *LRUCache) Delete(args request.CacheRequest) response.CacheResponse {
	return response.NewResponseFromMessage("METHODNOTAVAIL", 0)
}

func (cache *LRUCache) Flush(args request.CacheRequest) response.CacheResponse {
	return response.NewResponseFromMessage("METHODNOTAVAIL", 0)
}

func (cache *LRUCache) CountKeys(args request.CacheRequest) response.CacheResponse {
	return response.NewResponseFromMessage("METHODNOTAVAIL", 0)
}

func (cache *LRUCache) Ping(args request.CacheRequest) response.CacheResponse {
	return response.NewPingResponse()
}
