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

### Makefile Targets

| Command | Description |
|----------|-------------|
| `make run` | Run the CLI with default file (`urls.txt`) |
| `make run_json` | Run the CLI with default file (`urls.txt`) and outputs a json file (`results.json`)|
| `make test` | Run unit tests |


### 1. Run with a text file of URLs
```bash
make run --from-file ./urls.txt
```
Example file (`urls.txt`):
```
https://www.google.com
https://www.github.com
http://nonexistent.website
...
```

### 2. Run directly with flags
```bash
go run ./cmd --json --from-file ./urls.txt --concurrency 10
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
go run ./cmd --from-file ./urls.txt
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
+----------------------------------------------------------------+
| Ping Summary Report                                            |
+----------------------------+-----------------------------------+
| METRIC                     | VALUE                             |
+----------------------------+-----------------------------------+
| Total Successes            | 6                                |
| Total Slow Pings (> 1 sec) | 2                                 |
| Total Errors               | 4                                 |
| Average Latency (µs)       | 315526                            |
+----------------------------+-----------------------------------+
| Successful URLs            |                                   |
|                            | https://stackoverflow.com         |
|                            | https://www.wikipedia.org         |
|                            | https://www.bbc.co.uk             |
|                            | https://www.google.com            |
|                            | https://www.facebook.com          |
|                            | https://httpbin.org/status/404    |
+----------------------------+-----------------------------------+
| Slow URLs                  |                                   |
|                            | https://httpbin.org/status/500    |
|                            | https://httpbin.org/status/404    |
+----------------------------+-----------------------------------+
| Failed URLs                |                                   |
|                            | https://httpstat.us/503           |
|                            | https://httpstat.us/500           |
|                            | https://untrusted-root.badssl.com |
|                            | https://httpbin.org/delay/5       |
+----------------------------+-----------------------------------+
```

### JSON mode
```json
{
  "average_latency_microseconds": 360511,
  "failed_requests": 3,
  "failed_urls": [
    "https://self-signed.badssl.com",
    "https://untrusted-root.badssl.com",
    "https://httpbin.org/delay/5"
  ],
  "slow_requests": 1,
  "slow_urls": [
    "https://httpbin.org/status/404"
  ],
  "successful_urls": [
    "https://stackoverflow.com",
    "https://www.wikipedia.org",
    "https://www.bbc.co.uk",
    "https://www.google.com",
    "https://www.github.com",
    "https://www.nytimes.com",
    "https://www.amazon.com"
  ],
  "total_successes": 7
}

```

---

## License

MIT License © 2025 Danilo Barolo Martins de Lima
