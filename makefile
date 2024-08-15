# Variables
TASK = your-task
COMMIT = your-commit
ORIGIN = main

# Add and commit changes
add:
	git add . && git commit -m 'updated(${TASK}): ${COMMIT}'

# Push changes to origin
push:
	git push origin ${ORIGIN}

# Build the project
build:
	@go build -o bin/task-manager ./Delivery/main.go

# Run tests excluding the 'repos' folder
test:
	@go test $(shell go list ./... | grep -v '/repos') -v

# Run tests with coverage excluding the 'repos' folder
test-coverage:
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v '/repos')
	@go tool cover -func=coverage.out
	