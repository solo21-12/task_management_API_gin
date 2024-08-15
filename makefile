# Add and commit changes
add:
	git add . && git commit -m 'updated(${task}): ${commit}'

# Push changes to origin
push:
	git push origin ${ORIGIN}
# run the project
run:
	@go run ./Delivery/main.go


# Build the project
build:
	@go build -o bin/task-manager ./Delivery/main.go

# Run tests excluding the 'repos' folder
test:
	@go test $(shell go list ./... | grep -v '/Repositories') -v

# Run tests with coverage excluding the 'repos' folder
test-coverage:
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v '/Repositories')
	@go tool cover -func=coverage.out
	