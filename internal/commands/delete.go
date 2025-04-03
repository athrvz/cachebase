package commands

import "cachebase/internal/pkg/storage"

func Delete(key string) {
	mu := storage.Keylock.GetLock(key)
	mu.Lock()
	defer mu.Unlock()
	delete(storage.Cache, key)
}
