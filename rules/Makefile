.DEFAULT_GOAL := build

grammar:
	@echo "Generating grammar..."
	@antlr -Dlanguage=Go -visitor -o parser/__generated__ mindcarbon.g4

fmt:
	@echo "Formatting code..."
	@go fmt ./...

vet: fmt
	@echo "Running go vet..."
	@go vet ./...

build: grammar vet
	@echo "Building code..."
	@go build -o rules ./...