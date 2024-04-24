package base

import (
	"fmt"
	"testing"

	"github.com/bwlee13/gopherdb/storage/object"
	"github.com/bwlee13/gopherdb/storage/request"
	"github.com/bwlee13/gopherdb/storage/response"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {

	store := NewStore("LRU")
	store.BuildStore()
	// defer store.Close()
	if store == nil {
		t.Fatalf("Failed to create store instance")
	}

	expectedRes := response.CacheResponse{
		CacheObj: object.CacheObject{
			Key:   "",
			Value: interface{}(nil),
			TTL:   -1,
		},
		Status:  0,
		Message: "CACHE_MISS",
		Error:   "",
	}

	result := store.Execute("get", request.NewCacheRequest("Key1", "Val1", -1))
	fmt.Printf("Here is what res looks like: \n %v \n", result)
	assert.Equal(t, expectedRes, result)
}

func TestPutGet(t *testing.T) {
	store := NewStore("LRU")
	store.BuildStore()

	if store == nil {
		t.Fatalf("Failed to create store instance")
	}

	expectedResGet := response.CacheResponse{
		CacheObj: object.CacheObject{
			Key:   "",
			Value: "Val1",
			TTL:   -1,
		},
		Status:  1,
		Message: "OK",
		Error:   "",
	}

	resultPut := store.Execute("put", request.NewCacheRequest("Key1", "Val1", -1))
	resultGet := store.Execute("get", request.NewCacheRequest("Key1", "", -1))

	fmt.Printf("Here is what res PUT looks like: \n %v \n", resultPut)
	fmt.Printf("Here is what res GET looks like: \n %v \n", resultGet)

	assert.Equal(t, expectedResGet, resultGet)
}

func TestPutMessage(t *testing.T) {
	store := NewStore("LRU")
	store.BuildStore()

	if store == nil {
		t.Fatalf("Failed to create store instance")
	}

	expectedResPut := response.CacheResponse{
		CacheObj: object.CacheObject{
			Key:   "",
			Value: interface{}(nil),
			TTL:   -1,
		},
		Status:  1,
		Message: "STORED",
		Error:   "",
	}

	resultPut := store.Execute("put", request.NewCacheRequest("Key1", "Val1", -1))

	fmt.Printf("Here is what res PUT looks like: \n %v \n", resultPut)

	assert.Equal(t, expectedResPut.Message, resultPut.Message)
	assert.Equal(t, expectedResPut, resultPut)
}
