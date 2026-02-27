run:
	@ go run main.go

linter:
	@ golangci-lint run --timeout 5m