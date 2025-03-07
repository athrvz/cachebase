package commands

import (
	storage "cachebase/internal/pkg/storage"
	"time"
)

func Set(key string, value interface{}, options storage.SetOptions) (bool, storage.Record) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	record := storage.Record{
		Value: value,
	}
	if options.Expiry != nil {
		duration := time.Duration(*options.Expiry) * time.Millisecond
		ttl := time.Now().Add(duration)
		record.Expiry = &ttl
	}
	storage.Cache[key] = record
	return true, record
}
