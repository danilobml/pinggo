.PHONY: test build

test:
	go test ./...

run:
	go run ./cmd/... --from-file  ./urls.txt

run_json:
	go run ./cmd/... --json --from-file  ./urls.txt
