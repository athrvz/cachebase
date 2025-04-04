package commands

import (
	"cachebase/internal/pkg/storage"
)

// List Push Operation
func ListPush(key string, values []string) int {
	mu := storage.Keylock.GetLock(key)
	mu.Lock()
	defer mu.Unlock()

	// Get existing or create new
	record, exists := storage.Cache[key]
	var list []string
	if exists {
		list = record.Value.([]string)
	} else {
		list = []string{}
	}
	list = append(list, values...)
	storage.Cache[key] = storage.Record{
		Value:  list,
		Expiry: record.Expiry,
	}
	return len(list)
}

// List Pop Operation
func ListPop(key string) (string, bool) {
	mu := storage.Keylock.GetLock(key)
	mu.Lock()
	defer mu.Unlock()

	record, exists := storage.Cache[key]
	if !exists {
		return "", false
	}

	list, ok := record.Value.([]string)
	if !ok || len(list) == 0 {
		return "", false
	}

	value := list[len(list)-1]
	newList := list[:len(list)-1]

	if len(newList) == 0 {
		delete(storage.Cache, key)
	} else {
		storage.Cache[key] = storage.Record{
			Value:  newList,
			Expiry: record.Expiry,
		}
	}
	return value, true
}
