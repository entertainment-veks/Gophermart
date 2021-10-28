.PHONY: devrun
devrun:
	go run cmd/gophermart/main.go -a "localhost:8080" -d "host=localhost dbname=gophermart_db sslmode=disable user=postgres password=postgres" -r "accural.system.address"
.DEFAULT_GOAL := devrun

.PHONY: run
run:
	go run cmd/gophermart/main.go

.PHONY: build
build:
	go build -v -o Gophermart cmd/gophermart/main.go

.PHONY: test
	go test -v -race -timeout 30.0s ./...