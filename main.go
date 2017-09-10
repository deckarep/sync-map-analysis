package main

import (
	"fmt"
	"sync"
)

func main() {
	regularMapUsage()
	syncMapUsage()
}

func regularMapUsage() {
	fmt.Println("Regular threatsafe map test")
	fmt.Println("---------------------------")

	// Create the threadsafe map.
	reg := NewRegularStringMap()

	// Fetch an item that doesn't exist yet.
	result, ok := reg.Load("hello")
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	// Store an item in the map.
	reg.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	// Fetch the item we just stored.
	result, ok = reg.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result)
	}

	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Println()
}

func syncMapUsage() {
	fmt.Println("sync.Map test (Go 1.9+ only)")
	fmt.Println("----------------------------")

	// Create the threadsafe map.
	var sm sync.Map

	// Fetch an item that doesn't exist yet.
	result, ok := sm.Load("hello")
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	// Store an item in the map.
	sm.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	// Fetch the item we just stored.
	result, ok = sm.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result.(string))
	}

	fmt.Println("---------------------------")
}
