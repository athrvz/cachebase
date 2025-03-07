package commands

import "cachebase/internal/pkg/storage"

func Delete(key string) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.Cache, key)
}
