# PingGo

A lightweight, concurrent command-line tool written in Go that pings multiple websites simultaneously, measures response times, and reports availability and latency statistics in real time.

**Part of a series of mini-projects to get more knowledge of Go's features and ecosystem**

---

## Overview

**PingGo** demonstrates Go's concurrency model using goroutines, channels, contexts, and synchronization primitives. It takes a list of URLs, sends HTTP GET requests concurrently, and summarizes the results in either a table or JSON output.

---

## Installation

```bash
git clone https://github.com/danilobml/pinggo.git
cd pinggo
go mod tidy
```

---

## Usage

### 1. Run with a text file of URLs
```bash
make run
```
Example file (`test.txt`):
```
https://www.google.com
https://www.github.com
http://nonexistent.website
```

### 2. Run directly with flags
```bash
go run ./cmd --json --from-file ./test.txt --concurrency 10
```

### `--json`
Output results in **JSON format** instead of a table.  
Useful for automation or piping into tools like `jq`.

**Example:**
```bash
go run ./cmd --json --from-file urls.txt
```

---

### `--from-file`
Specify a text file containing one URL per line.  
Empty or commented lines (`#`) are ignored.

**Usage:**
```bash
go run ./cmd --from-file ./test.txt
```

---

### `--concurrency`
Set the **maximum number of concurrent pings** (default: `5`).  
Uses a buffered semaphore to limit active goroutines.

**Example:**
```bash
go run ./cmd --concurrency 10
```

### 3. Run tests
```bash
make test
```

---

## Example Output

### Table mode
```
URL                      STATUS  LATENCY     ERROR
------------------------------------------------------
https://google.com       200     123.3ms     -
https://github.com       200     187.2ms     -
http://nonexistent.xyz   -       -           dial tcp: lookup failed

SUMMARY:
Total: 3, Success: 2, Failed: 1, Avg Latency: 155.3ms
```

### JSON mode
```json
{
  "results": [
    {"url":"https://google.com","status":200,"latency_ms":123,"error":""},
    {"url":"https://github.com","status":200,"latency_ms":187,"error":""},
    {"url":"http://nonexistent.xyz","status":0,"latency_ms":0,"error":"lookup failed"}
  ],
  "summary": {"total":3,"success":2,"failed":1,"avg_latency_ms":155}
}
```

---

## Makefile Targets

| Command | Description |
|----------|-------------|
| `make run` | Run the CLI with default file (`urls.txt`) |
| `make run_json` | Run the CLI with default file (`urls.txt`) and outputs a json file (`results.json`)|
| `make test` | Run unit tests |

---

## License

MIT License © 2025 Your Name
