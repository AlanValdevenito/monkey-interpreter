.PHONY: help run run/file test test/scanner linter fmt vet tidy audit

# ---------- Helpers ----------

## help: Show available commands
help:
	@ echo "Available commands:"
	@ sed -n 's/^##//p' Makefile | column -t -s ':' | sed -e 's/^/ /'

# ---------- Run ----------

## run: Run the interpreter
run:
	@ echo "Running the interpreter..."
	@ go run main.go

## run/file: Run the interpreter with a specified file
run/file:
	@ echo "Running the interpreter with file: $(file)"
	@ go run main.go $(file)

# ---------- Tests ----------

## test: Run tests with coverage
test:
	@ echo "Running tests with coverage..."
	@ go test -v -cover ./...

## test/scanner: Run tests in the scanner package
test/scanner:
	@ echo "Running tests in the scanner package..."
	@ go test -v -cover ./scanner/...

# ---------- Quality control ----------

## linter: Run linter
linter:
	@ echo "Running linter..."
	@ golangci-lint run --timeout 5m

## fmt: Format code
fmt:
	@ echo "Formatting code..."
	@ go fmt ./...

## vet: Run go vet
vet:
	@ echo "Running go vet..."
	@ go vet ./... # Vet examines Go source code and reports suspicious constructs that are likely to be bugs

## tidy: Clean module dependencies
tidy:
	@ echo "Tidying module dependencies..."
	@ go mod tidy # Tidy ensures that our go.mod and go.sum files are accurate and clean

## audit: Run full quality checks
audit: tidy fmt vet linter test
	@ echo "Audit complete"