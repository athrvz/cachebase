package commands

import (
	"cachebase/internal/pkg/storage"
	"cachebase/internal/utilities"
)

func Get(key string) (interface{}, bool) {
	storage.Mutex.RLock()
	record, found := storage.Cache[key]
	storage.Mutex.RUnlock()
	if !found {
		return nil, false
	}

	// Check if the key has expired
	if utilities.IsExpired(record) {
		storage.Mutex.Lock()
		delete(storage.Cache, key)
		storage.Mutex.Unlock()
		return nil, false
	}
	return record.Value, true
}
