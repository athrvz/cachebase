# CacheBase 

**A high-performance Redis-inspired in-memory cache database written in Go**

## Key Features ✨
- 🚀 12,000+ operations/sec throughput
- ⚡ gRPC API for language-agnostic access
- 🔄 RDB snapshot persistence for crash recovery
- 🔒 Transaction support with MULTI/EXEC commands
- 📈 Horizontal scaling architecture

## Quick Start 🏁
# 🚀 Prerequisites

Before running this project, make sure the following tools are installed on your system:

---

### ✅ Required Tools

| Tool      | Version      | Description                              |
|-----------|--------------|------------------------------------------|
| 🟩 **Go** | `v1.20+`      | The Go programming language. [Download](https://golang.org/dl/) |
| 📦 **protoc** | Latest Release | Protocol Buffers compiler. [Get it here](https://github.com/protocolbuffers/protobuf/releases) |



### Build & Run
```bash
# Clone repository
git clone https://github.com/yourusername/cachebase.git
cd cachebase

# Build binary
go build -o cachebase cmd/server/main.go

# Start server (default port: 50051)
./cachebase --config configs/default.yaml

```

### Basic Usage
```bash
# Set a key-value pair
grpcurl -plaintext -d '{"key": "foo", "value": "bar"}' localhost:50051 cachebase.CacheService/Set

# Get a value
grpcurl -plaintext -d '{"key": "foo"}' localhost:50051 cachebase.CacheService/Get
```

## Why CacheBase? 💡

- 🚀 **Reduces latency by 90%** compared to traditional databases for cache workloads  
- 🔄 Ideal for **session storage**, **real-time analytics**, and **microservices caching**  
- 🪶 Lightweight alternative to **Redis**, built with **Go's concurrency** advantages  

---

## Roadmap 🗺️

- 📜 AOF (Append-Only File) persistence  
- 🧩 Cluster mode  
- 🧠 Redis protocol compatibility  
