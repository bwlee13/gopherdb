package cache

import (
	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/bwlee13/gopherdb/storage/response"
)

type Cache interface {
	Put(reqObj request.CacheRequest) response.CacheResponse
	Get(reqObj request.CacheRequest) response.CacheResponse
	Add(reqObj request.CacheRequest) response.CacheResponse
	Delete(reqObj request.CacheRequest) response.CacheResponse
	Flush(args request.CacheRequest) response.CacheResponse
	CountKeys(args request.CacheRequest) response.CacheResponse
	Ping(args request.CacheRequest) response.CacheResponse
}
