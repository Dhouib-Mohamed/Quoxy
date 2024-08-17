#!/bin/bash

rm -f coverage.out
rm -f test_output.txt
# Run tests and capture output
TEST_OUTPUT=$(ENV=TEST go test -v -cover -coverprofile=coverage.out ./internal/tests 2>&1)

# Save the test output to a file for debugging purposes
echo "$TEST_OUTPUT" > test_output.txt

# Extract coverage percentage
COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Define a minimum acceptable threshold (e.g., 80%)
SUCCESS_THRESHOLD=98.0

COVERAGE_THRESHOLD=65.0

# Extract success rate
SUCCESS_RATE=$(echo "$TEST_OUTPUT" | grep -c "ok")
TOTAL_TESTS=$(echo "$TEST_OUTPUT" | grep -E "(FAIL|ok)" | wc -l)

# Calculate success percentage
if [ "$TOTAL_TESTS" -ne 0 ]; then
    SUCCESS_PERCENTAGE=$(echo "scale=2; ($SUCCESS_RATE / $TOTAL_TESTS) * 100" | bc)
else
    SUCCESS_PERCENTAGE=0
fi

echo "Coverage: $COVERAGE%"
echo "Success Rate: $SUCCESS_PERCENTAGE%"

# Extract failed tests and their causes
FAILED_TESTS=$(echo "$TEST_OUTPUT" | grep -E "FAIL: |--- FAIL:")

if [ -n "$FAILED_TESTS" ]; then
    echo "Failed Tests:"
    echo "$FAILED_TESTS"
else
    echo "No tests failed."
fi

# Check if the success percentage and coverage meet the threshold
if (( $(echo "$COVERAGE < $COVERAGE_THRESHOLD" | bc -l) )); then
    echo "Insufficient coverage."
    exit 1
fi

if (( $(echo "$SUCCESS_PERCENTAGE < $SUCCESS_THRESHOLD" | bc -l) )); then
    echo "Many Tests Failed"
    exit 1
fi

echo "Tests passed with sufficient coverage."
exit 0
