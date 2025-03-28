package transaction

import "sync"

type KeyLock struct {
	Mutex sync.Mutex
	locks map[string]*sync.Mutex
}

func (kl *KeyLock) Lock(key string) {
	kl.Mutex.Lock()
	if _, found := kl.locks[key]; !found {
		kl.locks[key] = &sync.Mutex{}
	}
	m := kl.locks[key]
	kl.Mutex.Unlock()
	m.Lock()
}

func (kl *KeyLock) Unlock(key string) {
	kl.Mutex.Lock()
	m := kl.locks[key]
	kl.Mutex.Unlock()
	m.Unlock()
}
