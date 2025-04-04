package commands

import (
	storage "cachebase/internal/pkg/storage"
	"sync/atomic"
	"time"
)

func Set(key string, value interface{}, options storage.SetOptions) (bool, storage.Record) {
	atomic.AddUint64(&storage.WriteOperations, uint64(1))
	// storage.Mutex.Lock()
	mu := storage.Keylock.GetLock(key)
	mu.Lock()
	defer mu.Unlock()

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
