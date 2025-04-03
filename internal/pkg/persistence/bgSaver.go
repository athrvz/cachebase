package persistence

import (
	"cachebase/internal/config"
	"cachebase/internal/pkg/storage"
	"fmt"
	"sync/atomic"
	"time"
)

func StartRDBSaver(config config.RDBConfig, stopChan <-chan struct{}) {
	for _, interval := range config.SaveIntervals {
		go func(secs, changes int) {
			ticker := time.NewTicker(time.Duration(secs) * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-ticker.C:
					if atomic.LoadUint64(&storage.WriteOperations) >= uint64(changes) {
						fmt.Println("writeOps: ", storage.WriteOperations)
						ScheduleRDBSave(config)
						atomic.StoreUint64(&storage.WriteOperations, 0)
					}
				case <-stopChan:
					return
				}
			}
		}(interval.Seconds, interval.Changes)
	}
}
