package transaction

import (
	"cachebase/internal/commands"
	"cachebase/internal/pkg/storage"

	"github.com/google/uuid"
)

// MULTI
func (tm *TransactionManager) Begin(cliendID string) {
	trnx := &Transaction{
		ID:       uuid.New().String(),
		Commands: make(chan *Command, 100), // Buffered Channel
		ReadSet:  make(map[string]struct{}),
		WriteSet: make(map[string]struct{}),
	}
	trnx.Active.Store(true)
	tm.Transactions.Store(cliendID, trnx)

	// Goroutine to process cmds
	go func() {
		for cmd := range trnx.Commands {
			trnx.Wg.Add(1)
			go tm.processCommand(trnx, cmd)
		}
	}()
}

func isWriteCommand(op string) bool {
	switch op {
	case "SET", "DELETE", "LPUSH", "RPUSH", "SADD", "SREM":
		return true
	default:
		return false
	}
}

func (tm *TransactionManager) processCommand(tx *Transaction, cmd *Command) {
	defer tx.Wg.Done()

	// Conflict detection
	if isWriteCommand(cmd.Op) {
		for _, arg := range cmd.Args {
			if key, ok := arg.(string); ok {
				tm.GlobalLock.Lock(key) // Acquire key-level lock
				tx.WriteSet[key] = struct{}{}
			}
		}
	}

	// Execute command (eg: SET, GET)
	switch cmd.Op {
	case "SET":
		key := cmd.Args[0].(string)
		value := cmd.Args[1].(string)
		opts := storage.SetOptions{Expiry: nil}
		commands.Set(key, value, opts)
		cmd.Result <- "OK"

	case "GET":
		key := cmd.Args[0].(string)
		value, found := commands.Get(key)
		if !found {
			cmd.Result <- "NOT_FOUND"
			break
		}
		cmd.Result <- value.(string)
	}

	// Release locks after write
	if isWriteCommand(cmd.Op) {
		for _, arg := range cmd.Args {
			if key, ok := arg.(string); ok {
				tm.GlobalLock.Unlock(key)
			}
		}
	}
}

// EXEC (commit trnx)
func (tm *TransactionManager) Commit(cliendID string) ([]any, error) {
	rawTx, _ := tm.Transactions.Load(cliendID)
	tx := rawTx.(*Transaction)
	close(tx.Commands) // stop accepting new commands
	tx.Wg.Wait()       // Wait for all cmds to finish

	// Collect results
	var results []any
	for cmd := range tx.Commands {
		select {
		case res := <-cmd.Result:
			results = append(results, res)
		case err := <-cmd.Err:
			return nil, err
		}
	}
	tm.Transactions.Delete(cliendID)
	return results, nil
}
