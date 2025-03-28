# Transaction Design for Cachebase

## 1. Key Concepts

- **Atomicity**: All commands in a transaction succeed or fail together.
- **Isolation**: Transactions are serialized (no interleaving with other operations).
- **No Rollbacks**: Like Redis, we won’t support rollbacks (simplicity over complexity).

## 2. Commands

| Command   | Description                                |
|-----------|--------------------------------------------|
| **MULTI** | Start a transaction block.                |
| **EXEC**  | Execute all commands in the block.        |
| **DISCARD** | Abort the transaction.                   |
