.PHONY: build
build:
	go build -o app cmd/main.go

.PHONY: run
run:
	go run cmd/main.go -level=debug

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: imports
imports:
	find . -name \*.go -exec goimports -w {} \;

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm app
