package persistence

import (
	"cachebase/internal/config"
	"cachebase/internal/pkg/storage"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	activeSaves    atomic.Int32
	lastSaveStatus atomic.Value
	mu             sync.RWMutex
)

type saveStatus struct {
	Time  time.Time
	Error error // Can be nil
}

// Non-blocking Save
func ScheduleRDBSave(config config.RDBConfig) error {
	// if !config.BackgroundSave {
	// 	return saveRDBBlocking(config) // Original Blocking save
	// }
	if activeSaves.Load() >= int32(config.MaxConcurrentSaves) {
		return errors.New("too many concurrent saves")
	}
	activeSaves.Add(1)
	defer activeSaves.Add(-1)

	// 1. Create isolated snapshot
	mu.RLock()
	snapshot := make(map[string]storage.Record, len(storage.Cache))
	for k, v := range storage.Cache {
		snapshot[k] = v // Shallow copy
	}
	mu.RUnlock()

	// 2. Process in background
	go func() {
		start := time.Now()
		err := saveSanpshot(config, snapshot)
		status := &saveStatus{Time: start, Error: err}
		lastSaveStatus.Store(status)

		if err != nil {
			log.Printf("Background save failed: %v", err)
		}
	}()
	return nil
}

// Actual disk write (with compression)
func saveSanpshot(config config.RDBConfig, snapshot map[string]storage.Record) error {
	// 1. Create temp file
	dynamicFileName := config.FilePath + config.TempFilePattern + config.CdbFileExtension
	tempFile, err := os.Create(dynamicFileName)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// 2. Compression
	gz := gzip.NewWriter(tempFile)
	defer gz.Close()

	enc := gob.NewEncoder(gz)
	if err := enc.Encode(snapshot); err != nil {
		return err
	}
	return nil
}

func GetLastSavedStatus() (time.Time, error) {
	status, ok := lastSaveStatus.Load().(*saveStatus)
	if !ok || status == nil {
		return time.Time{}, errors.New("no save recorded")
	}
	return status.Time, status.Error
}

func LoadRDB(config config.RDBConfig) error {
	dynamicFileName := config.FilePath + config.TempFilePattern + config.CdbFileExtension
	// Check if RDB file exists
	if _, err := os.Stat(dynamicFileName); os.IsNotExist(err) {
		return fmt.Errorf("cdb file not found: %s", dynamicFileName)
	}

	// Open file with compression detection
	file, err := os.Open(dynamicFileName)
	if err != nil {
		return fmt.Errorf("error opening CDB file: %w", err)
	}
	defer file.Close()

	// Wrap the decompression
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("gzip decompression failed: %v", err)
	}
	defer gzReader.Close()

	// Atomic load with full lock
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()

	// Decode using gob
	decoder := gob.NewDecoder(gzReader)
	if err := decoder.Decode(&storage.Cache); err != nil {
		storage.Cache = make(map[string]storage.Record)
		return fmt.Errorf("cdb decoding failed: %v", err)
	}

	return nil
}
