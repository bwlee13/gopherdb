package request

import "github.com/bwlee13/gopherdb/storage/object"

type CacheRequest struct {
	CacheObj object.CacheObject
}

func NewCacheRequest(key string, value interface{}, ttl int64) CacheRequest {
	return CacheRequest{
		CacheObj: object.NewCacheObjFromParams(key, value, ttl),
	}
}

func NewEmptyCacheRequest() CacheRequest {
	return CacheRequest{
		CacheObj: object.NewEmptyCacheObj(),
	}
}
