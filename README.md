Go Lab 🧪

A Go learning and practice repository: small applications, concurrency examples, coding challenges, and exercises in structs, pointers, and API design.

## Repository structure

| Directory | Description |
|-----------|-------------|
| **taxation** | Tax-included price calculator: reads prices from a file, applies multiple tax rates concurrently, writes JSON results. |
| **bankApp** | Bank application with file-backed storage and a presentation layer. |
| **profitCalculator** | Profit calculation utility. |
| **Challenges** | Coding challenges (NeetCode, array problems, strings, recursion, etc.). |
| **concurrencyProblems** | Concurrency patterns: contexts, pipelines, workers, fan-out, race conditions, wait groups. |
| **Pointers** | Pointer practice and game-stats–style examples. |
| **StructsPractice** | Struct usage: notes, animals, users. |
| **apiDesign** | Simple API and API design examples. |
| **NamesProject** | User/structs naming and organization practice. |

## Taxation app (main runnable project)

Computes tax-included prices for a list of input prices at several tax rates (0%, 7%, 10%, 15%), using concurrent jobs and file I/O.

### Prerequisites

- Go 1.25+

### Run

From the `taxation` directory:

```bash
cd taxation
go run tax.go
```

### Input / output

- **Input:** `taxation/prices.txt` — one price per line (e.g. `12`, `9.76`, `3.23`).
- **Output:** `result_0.json`, `result_70.json`, `result_10.json`, `result_15.json` in the same directory (tax rate encoded in filename; 0 = 0%, 70 = 7%, etc.).

### Packages

- **taxation** — `main`; spawns one goroutine per tax rate and coordinates completion/errors via channels.
- **taxation/filemanager** — reads lines from the input file and writes result structs as JSON.
- **taxation/prices** — `TaxIncludedPrice` job: loads prices, computes tax-included amounts, writes result via the file manager.
- **taxation/conversion** — converts slices of strings to `[]float64` for price parsing.

## Other projects

- **Challenges:** Run individual challenge mains (e.g. `go run ./Challenges/...`) as needed.
- **concurrencyProblems, Pointers, StructsPractice, apiDesign, NamesProject:** Same idea — navigate to the relevant directory and run the main package or example (e.g. `go run .` or `go run main.go`).


