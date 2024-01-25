# Tests

### test file requirements?
Filename ends with _test.go
Imports `testing` package 

### 3 types of function supported by go in test files?
Test*
Benchmark*
Example*


### Is t.Error stopping the execution of test case?
No, all test case will continue to run. To stop the execution, run t.Fatal

### Why do we need external test packages?
Mostly to create integration tests that cannot be created in the same package due to cyclic dependencies. 
These packages should be name package_test

### How to list all files that are included in executable? All test and external test?
- app files `go list -f={{.GoFiles}} <package>`
- test files `go list -f={{.TestGoFiles}} <package>`
- external test files `go list -f={{.XTestGoFiles}} <package>`