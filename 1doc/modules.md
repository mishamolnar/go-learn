# Modules
### What is go package?
Collection of go files in the same directory compiled together

### What is go module?
Collection of related go packages that are released together. Usually go repo contains 1 module in the root of the repo. 
File go.mod declares module path 

### What does `go install` command?
Builds executable binary, and installs it in ` $GOPATH/bin`

### Where the dependencies files are downloaded to?
They are downloaded to `$GOPATH/pkg/mod` and marked as read only

### How can I remove external dependencies files?
`go clean -modcache`

### 