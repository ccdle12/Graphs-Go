.PHONY: check main

default: main

# Run main.go
main:
	@go run main.go

# Run Tests.
check:
	@go test -v ./...
