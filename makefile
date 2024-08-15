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

# Run tests
test:
	@go test ./tests/... -v

# Run tests with coverage
test-coverage:
	@go test -coverprofile=coverage.out ./tests/...
	@go tool cover -func=coverage.out
