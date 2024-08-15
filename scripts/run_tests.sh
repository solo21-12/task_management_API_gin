#!/bin/bash

# Exit on error
set -e

# List of packages to exclude (adjust as needed)
exclude_packages=(
    "task_seven/tests/task_tests/repo"
    "task_seven/tests/user_tests/repo"
    "task_seven/tests/constants"
)

# Find all packages with .go files in the tests directory
all_packages=$(go list ../tests/...)

# Filter out the excluded packages
for exclude in "${exclude_packages[@]}"; do
    all_packages=$(echo "$all_packages" | grep -v "^${exclude}$")
done

# Convert to a space-separated list
test_packages_list=$(echo "$all_packages" | tr '\n' ' ')

# Print the list of packages
echo "Test packages to run:"
echo $test_packages_list

# Run tests for the filtered packages
echo "Running tests..."
go test $test_packages_list
