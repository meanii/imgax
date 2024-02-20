BINARY=mai
BIN_DIR=bin

dep:
	@echo "Installing dependencies"
	go mod tidy
	@echo "Dependencies installed"

compile:
	@echo "Compiling"
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) cmd/main.go
	@echo "Compiled"

build: dep compile
	@echo "Build complete"

run:
	@echo "Running"
	./$(BIN_DIR)/$(BINARY)

dev:
	@echo "Running in dev mode..."
	@go run main.go

watch:
	@nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
