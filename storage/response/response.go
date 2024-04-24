package response

import "github.com/bwlee13/gopherdb/storage/object"

const (
	INVALID_COMMAND_ERR = "INVALID_COMMAND_ERR"
)

type CacheResponse struct {
	CacheObj object.CacheObject
	// Status of result (0: Fail, 1: Success)
	Status int32
	// text data to be returned to client
	Message string
	// Error message to be returned to client
	Error string
}

func NewResponseFromValue(value interface{}) CacheResponse {
	return CacheResponse{
		CacheObj: object.NewCacheObjFromValue(value),
		Status:   1,
		Message:  "OK",
		Error:    "",
	}

}

func NewCacheMissResponse() CacheResponse {
	return CacheResponse{
		CacheObj: object.NewEmptyCacheObj(),
		Status:   0,
		Message:  "CACHE_MISS",
		Error:    "",
	}
}

func NewPingResponse() CacheResponse {
	return CacheResponse{
		CacheObj: object.NewEmptyCacheObj(),
		Status:   1,
		Message:  "PONG!",
		Error:    "",
	}
}

func NewResponseFromMessage(msg string, status int32) CacheResponse {
	return CacheResponse{
		CacheObj: object.NewEmptyCacheObj(),
		Status:   status,
		Message:  msg,
		Error:    "",
	}
}
