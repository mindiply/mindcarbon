.DEFAULT_GOAL := build

fmt:
	@echo "Formatting code..."
	@go fmt ./...

vet: fmt
	@echo "Running go vet..."
	@go vet ./...

build: grammar vet
	@echo "Building code..."
	@go build -o rules ./...

test:
	@echo "Running tests..."
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Tests completed. Coverage report generated at coverage.html"

grammar:
	@echo "Generating grammar..."
	@antlr -Dlanguage=Go -visitor -o parser/__generated__ mindcarbon.g4

jgrammar:
	@echo "Generating Java grammar..."
	@antlr -Dlanguage=Java -visitor -o parser/__generated__ mindcarbon.g4
	cd parser/__generated__ && \
		javac *.java && \
		cd ../..
