package commands

import (
	"cachebase/internal/pkg/storage"
)

func ListPush(key string, values []string) int {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

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

func ListPop(key string) (string, bool) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	record, exists := storage.Cache[key]
	if !exists {
		return "", false
	}

	list, ok := record.Value.([]string)
	if !ok || len(list) == 0 {
		return "", false
	}

	value := list[len(list)-1]
	list = list[:len(list)-1]

	if len(list) == 0 {
		delete(storage.Cache, key)
	} else {
		storage.Cache[key] = storage.Record{
			Value:  list,
			Expiry: record.Expiry,
		}
	}
	return value, true
}
