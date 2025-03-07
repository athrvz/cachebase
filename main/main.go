package main

import (
	"cachebase/internal/commands"
	"cachebase/internal/pkg/storage"
	"fmt"
	"time"
)

func main() {
	options := storage.SetOptions{
		Expiry: nil,
	}
	commands.Set("foo", "bar", options)
	value, found := commands.Get("foo")
	if found {
		fmt.Println("foo: ", value)
	}
	ttl := 2
	options.Expiry = &ttl
	commands.Set("foo", "barWithExp", options)
	value, found = commands.Get("foo")
	if found {
		fmt.Println("foo: ", value)
	}

	time.Sleep(3 * time.Millisecond)
	value, found = commands.Get("foo")
	if !found {
		fmt.Println("Not found")
	}
}
