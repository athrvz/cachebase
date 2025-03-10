package commands

import (
	"cachebase/internal/pkg/storage"
	"time"
)

func Get(key string) (interface{}, bool) {
	storage.Mutex.RLock()
	record, found := storage.Cache[key]
	storage.Mutex.RUnlock()
	if !found {
		return nil, false
	}

	// Check if the key has expired
	if record.Expiry != nil && time.Now().After(*record.Expiry) {
		storage.Mutex.Lock()
		delete(storage.Cache, key)
		storage.Mutex.Unlock()
		return nil, false
	}

	return record.Value, true
}
