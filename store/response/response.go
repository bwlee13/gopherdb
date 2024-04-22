package response

import "github.com/bwlee13/gopherdb/store/object"

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
