package test

import (
	"cachebase/internal/commands"
	"cachebase/internal/pkg/storage"
	"testing"
	"time"
)

// TestDelete tests the Delete function.
func TestDelete(t *testing.T) {
	key := "foo"
	value := "bar"
	opts := storage.SetOptions{
		Expiry: nil,
	}
	// Set a value first
	commands.Set(key, value, opts)

	// Test Delete
	commands.Delete(key)
	_, exists := storage.Cache[key]
	if exists {
		t.Errorf("Expected key %s to be deleted", key)
	}
}

func TestSet(t *testing.T) {
	key, value, ttl := "foo", "bar", 2000
	opts := storage.SetOptions{
		Expiry: &ttl,
	}
	commands.Set(key, value, opts)
	record, exists := storage.Cache[key]
	if !exists {
		t.Errorf("Expected key %s to exists", key)
	}
	if record.Value != value {
		t.Errorf("Expected value: %s, got %v", value, record.Value)
	}
	// checking for ttl
	time.Sleep((time.Duration)(ttl/1000+1) * time.Second)
	_, exists = commands.Get(key)
	if exists {
		t.Errorf("Expected key to be expired")
	}
}

func TestGet(t *testing.T) {
	key, value := "foo", "bar"
	defer commands.Delete(key)
	opts := storage.SetOptions{
		Expiry: nil,
	}
	commands.Set(key, value, opts)

	resValue, exists := commands.Get(key)
	if !exists {
		t.Errorf("Expected Key: %s to exists", key)
	}
	if resValue != value {
		t.Errorf("Expected value: %s, got %v", value, resValue)
	}

	_, exists = commands.Get("nonexistent")
	if exists {
		t.Errorf("Expected key to not exist")
	}
}

func TestListPush(t *testing.T) {
	key, values := "list", []string{"one", "two"}
	defer commands.Delete(key)

	length := commands.ListPush(key, values)

	if length != len(values) {
		t.Errorf("Expected list length: %d, got %d", len(values), length)
	}

	record, exists := storage.Cache[key]
	if !exists {
		t.Errorf("Expected key %s to exist", key)
	}

	list, ok := record.Value.([]string)
	if !ok {
		t.Errorf("Expected value to be of type []string")
	}

	for i, v := range list {
		if v != values[i] {
			t.Errorf("Expected value %s at index %d, got %s", values[i], i, v)
		}
	}
}

func TestListPop(t *testing.T) {
	key, values := "list", []string{"one", "two", "three"}

	commands.ListPush(key, values)

	for i := len(values) - 1; i >= 0; i-- {
		value, exists := commands.ListPop(key)
		if !exists {
			t.Errorf("Expected value %s, got %s", values[i], value)
		}
	}

	_, exists := commands.ListPop(key)
	if exists {
		t.Errorf("Expected list to be empty")
	}
}
