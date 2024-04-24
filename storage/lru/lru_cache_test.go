package lru

import (
	"testing"

	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	cache := NewLRU()
	cache.Size = int32(2)

	cache.Put(request.NewCacheRequest("K1", "V1", -1))
	cache.Put(request.NewCacheRequest("K2", "V2", -1))
	// Head -> V2 -> V1

	cache.Put(request.NewCacheRequest("K3", "V3", -1)) // Evict K1: V1
	// Head -> V3 -> V2
	v1 := cache.Get(request.NewCacheRequest("K1", "", -1)) // Cache Miss
	assert.Equal(t, "CACHE_MISS", v1.Message)

	// Should move V2 to front of Cache as MRU
	v2 := cache.Get(request.NewCacheRequest("K2", "", -1))
	assert.Equal(t, "V2", v2.CacheObj.Value)
	// Head -> V2 -> V3
	cache.Put(request.NewCacheRequest("K4", "V4", -1)) // Evict K3: V3
	// Head -> V4 -> V2

	v3 := cache.Get(request.NewCacheRequest("K3", "", -1)) // Cache Miss
	assert.Equal(t, "CACHE_MISS", v3.Message)

}
