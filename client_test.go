package main

import (
	"fmt"
	"testing"

	gopherdb "github.com/bwlee13/gopherdb/notgopherdb"
)

// TestCacheSetAndGet tests the functionality of the Set and Get methods of the Cache
func TestCacheSetAndGet(t *testing.T) {
	go StartServer()
	cache, err := gopherdb.NewCache("localhost:42069")
	if err != nil {
		t.Fatalf("Failed to connect to cache server: %v", err)
	}
	defer cache.Close()

	// Test data
	testKey := "hello"
	testValue := "world"

	// Set a value in the cache
	if err := cache.Set(testKey, testValue); err != nil {
		t.Errorf("Failed to set value: %v", err)
	}

	// Retrieve a value from the cache
	value, err := cache.Get(testKey)
	if err != nil {
		t.Errorf("Failed to get value: %v", err)
	}

	fmt.Println("Value: ", value)
	// cache.Close()
	// t.Cleanup(func() { os.Exit(0) })

	// Check if the retrieved value matches the set value
	if value != testValue {
		t.Errorf("Expected value '%s', got '%s'", testValue, value)
	}
}
