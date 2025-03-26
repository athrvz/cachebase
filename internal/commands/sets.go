package commands

import (
	"cachebase/internal/pkg/storage"
)

func SetAdd(key string, values []string, options storage.SetOptions) int {
	// Get existing set
	resValue, found := Get(key)
	var set map[string]bool
	if found {
		set = resValue.(map[string]bool)
	} else {
		set = make(map[string]bool)
	}
	// Adding values to set
	added := 0
	for _, value := range values {
		if _, found = set[value]; !found {
			set[value] = true
			added++
		}
	}
	// Update record
	done, _ := Set(key, set, options)
	if done {
		return added
	}
	return 0
}

func SetRemove(key string, values []string) int {
	storage.Mutex.Lock()
	// Get existing record
	record, found := storage.Cache[key]
	if !found {
		return 0
	}

	set, ok := record.Value.(map[string]bool)
	if !ok {
		return 0
	}
	// Remove from set
	removed := 0
	for _, value := range values {
		if _, found := set[value]; found {
			delete(set, value)
			removed++
		}
	}

	if len(set) == 0 {
		storage.Mutex.Unlock()
		Delete(key)
	} else {
		storage.Cache[key] = storage.Record{
			Value:  set,
			Expiry: record.Expiry,
		}
		storage.Mutex.Unlock()
	}
	return removed
}

func SetIsMember(key, value string) bool {

	resultSet, found := Get(key)
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()
	if !found {
		return false
	}
	set, ok := resultSet.(map[string]bool)
	if !ok {
		return false
	}
	_, found = set[value]
	return found
}

func SetMembers(key string) ([]string, bool) {
	storage.Mutex.RLock()
	defer storage.Mutex.RUnlock()

	resValue, found := Get(key)
	if !found {
		return nil, false
	}
	set, ok := resValue.(map[string]bool)
	if !ok {
		return nil, false
	}
	members := make([]string, 0, len(set))
	for member := range set {
		members = append(members, member)
	}
	return members, true
}
