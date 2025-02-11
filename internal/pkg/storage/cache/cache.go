package cache

import "time"

type SetOptions struct {
  Expiry *int
}

type Record struct {
  value string
  expiryTS *time.Time
}

var cache = map[string]Record{}

func Set(key, val string, options SetOptions) {
  record := Record{
          value: val
  }
  if options.Expiry != nil {
    duration := time.Duration(*options.Expiry) * time.Millisecond
    expiryTS := time.Now().Add(duration)
    record.expiryTS = &expiryTS
  }
  cache[key] = record
}

func Get(key string) (string, bool) {
  val, ok := cache[key]
  if !ok || (val.expiryTS != nil && time.Now().After(*val.expiryTS)) {
    return "", false
  }
  return val.value, true
}
