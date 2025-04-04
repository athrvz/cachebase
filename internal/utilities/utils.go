package utilities

import (
	"cachebase/internal/pkg/storage"
	"time"
)

func IsExpired(record storage.Record) bool {
	return record.Expiry != nil && time.Now().After(*record.Expiry)
}

func GetCachedKeysCount() int {
	return len(storage.Cache)
}
