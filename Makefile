run:
	@ go run main.go

linter:
	@ golangci-lint run --timeout 5m

test:
	@ go test -v -cover ./...

test/scanner:
	@ go test -v -cover ./scanner/...