package storage

import (
	"sync"
	"time"
)

type Record struct {
	Value  interface{}
	Expiry *time.Time
}

type SetOptions struct {
	Expiry *int
}

var (
	Cache = make(map[string]Record)
	Mutex = &sync.RWMutex{}
)
