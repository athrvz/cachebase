package transaction

import (
	"sync"
	"sync/atomic"
)

type Command struct {
	Op     string        // SET, GET, etc
	Args   []interface{} // Command args
	Result chan any      // Channel for cmd result
	Err    chan error
}

type Transaction struct {
	ID       string
	Commands chan *Command // Buffered channel for command pipeline
	Wg       sync.WaitGroup
	Active   atomic.Bool
	Mutex    sync.Mutex
	ReadSet  map[string]struct{} // Track read keys
	WriteSet map[string]struct{} // Track write keys
}

type TransactionManager struct {
	Transactions sync.Map // Concurrent map [clientID]*Transaction
	GlobalLock   *KeyLock
}
