package storage

import (
	"sync"
)

type KeyLock struct {
	locks map[string]*sync.RWMutex
}

func (kl *KeyLock) GetLock(key string) *sync.RWMutex {
	Mutex.Lock()
	defer Mutex.Unlock()

	// Ensure there is no lock on the key
	if _, alreadyLocked := kl.locks[key]; !alreadyLocked {
		kl.locks[key] = &sync.RWMutex{}
	}
	return kl.locks[key]
}
