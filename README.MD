# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

You can download the Go Programming Language at: https://go.dev/doc/install

You can also view the Go documentation at: https://go.dev/doc/

## Repo Format
There is one folder called grade-calculator that contains the files:
- go.mod: Defines module properties, dependencies, and the version of Go.
- grade_calculator.go: Implements structs, data types, and core functions for grade calculation.
- grade_calculator_test.go: Contains unit tests to verify correct functionality of the grade calculator.

## Testing
Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
