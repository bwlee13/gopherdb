package main

import (
	"fmt"
	"github.com/bwlee13/gopherdb/gopherdb"  // Adjust the path according to your module setup
	"os"
)

func main() {
	// Initialize the cache connection
	cache, err := gopherdb.NewCache("localhost:42069")
	if err != nil {
		fmt.Println("Failed to connect to cache server:", err)
		os.Exit(1)
	}
	defer cache.Close()

	// Set a value in the cache
	if err := cache.Set("hello", "world"); err != nil {
		fmt.Println("Failed to set value:", err)
		return
	}

	// Retrieve a value from the cache
	value, err := cache.Get("hello")
	if err != nil {
		fmt.Println("Failed to get value:", err)
		return
	}
	fmt.Println("Value:", value)
}
