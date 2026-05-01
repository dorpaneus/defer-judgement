# gopher-garbage 🐹

My chaotic Go learning repo. Algorithms, random scripts, half-baked APIs, and future K8s operators.  
Contributions welcome? Nah, this is my dumpster fire. 🔥

## Structure
- `cmd/` → runnable things
- `internal/algorithms/` → DS&A practice
- `internal/scripts/` → small utilities

## Running examples
```bash
go run ./cmd/hello-world
go test ./internal/algorithms/...

This keeps things discoverable while staying flexible. As you grow, you can extract popular pieces into their own repos.

Which name do you like most? Want me to tweak the structure for more emphasis on algorithms or operators? Or generate a sample `main.go` for one of the cmd folders? Just say the word!
