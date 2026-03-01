run:
	@ go run main.go

run/file:
	@ go run main.go $(file)

linter:
	@ golangci-lint run --timeout 5m

test:
	@ go test -v -cover ./...

test/scanner:
	@ go test -v -cover ./scanner/...