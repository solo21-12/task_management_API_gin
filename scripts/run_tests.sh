#!/bin/bash

# Exit on error
set -e

# Change to the directory where the script is located
cd "$(dirname "$0")/.."

# Module path
module_path="github.com/solo21-12/A2SV_back_end_track/task_seven"

# List of packages to exclude (full path format)
exclude_packages=(
    "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/task_tests/repo"
    "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/user_tests/repo"
    "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/tests/constants"
)

# Find all packages with .go files in the tests directory
all_packages=$(go list ./tests/...)

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
if [ -n "$test_packages_list" ]; then
    echo "Running tests..."
    go test $test_packages_list
else
    echo "No test packages to run."
fi


