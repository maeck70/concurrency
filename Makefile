# Go-related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get $(MAIN_FILE)
BINARY_NAME=concurrency  # Name of your executable
MAIN_FILE=./cmd/.  # Path to your main file

generate:
	clear
	templ generate ./content

# Builds the executable
build: # generate
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Cleans up build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Runs tests
test:
	$(GOTEST) -v ./...

# Installs dependencies
install:
	$(GOGET) -u ./...

# Runs the application
run: build
	./$(BINARY_NAME)

# Runs the application with live reloading on file changes
run/live:
	go run github.com/cosmtrek/air@latest \
		--build="make build" --bin="$(BINARY_NAME)" \
		--delay=100ms --exclude_dir=node_modules

# Watches for file changes and re-runs tests
watch:
	gow -c make test

# Additional targets for static analysis, linting, etc. (optional)
lint:
	# ...

staticcheck:
	# ...

# PHONY targets to avoid conflicts with files of the same name
.PHONY: build clean test install run run/live watch lint staticcheck
