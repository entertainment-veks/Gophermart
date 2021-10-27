.PHONY: run
run:
	go run cmd/gophermart/main.go -a "localhost:8080" -d "databaseURI" -r "accural.system.address"
.DEFAULT_GOAL := run

.PHONY: build
build:
	go build -v -o Gophermart cmd/gophermart/main.go

.PHONY: test
	go test -v -race -timeout 30.0s ./...
