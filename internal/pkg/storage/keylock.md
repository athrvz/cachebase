# Key-Level Locking for a Concurrent Map 🗝️🔒

This implementation demonstrates a safe, concurrent map in Go where each **key** in the map is protected by its own lock. This allows **efficient concurrent access** to different keys while ensuring exclusive access for writes and concurrent reads for each key.

## Problem Overview 🤔

In Go, the built-in `map` type is **not thread-safe** by default. If multiple goroutines attempt to access or modify a map concurrently, it can result in a **data race**, leading to unpredictable behavior or crashes. To solve this, we need to ensure that:

- **Multiple goroutines** can **read** values from different keys simultaneously.
- **Exclusive access** for **writes** to each key, meaning no other goroutine can read or write that key while a write is in progress.

## Solution Overview 🛠️

To solve this problem, we use key-level locking, where each key has its own `sync.RWMutex`. This allows for **fine-grained control** over access to individual map keys.

### Components:
- **`sync.RWMutex`**: A read-write lock used to protect each key in the map.
- **`sync.Mutex`**: A global mutex used to protect the map of locks, ensuring only one goroutine can modify the lock map at a time.

---
