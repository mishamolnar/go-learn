# Questions fundamentals

### 1. Types in go?

- value types: number, booleans, strings 
- reference types: slice, pointer, map, chanel, function
- aggregate types: array and struct 

### 2. Default values for all types
- 0 for numbers, false for booleans, "" for strings
- nil for references
- zero value for each element of aggregate type

### 3. Pointers terminology?
- &x - address of x
- *int - pointer to int
- *p - p contains the address of x


### 4. Why do we need new(Type) function?
In order to get a pointer to the unnamed variable. This is no more than syntax sugar:
These functions are identical:
```go
package main

func newInt() *int {
	return new(int)
}

func newInt() *int {
	var dummyName int
	return &dummyName
}
```
new declaration does not determine if variable will be in the heap or stack

### 5. Named type syntax

`type name underlying-type`
```go
package main
// Named type example
type Celsius float64
type Fahrenheit float64
```

### 6. Named typed conversion. What will this command print? 

```go
package main

import "fmt"

type Kilogram float64
type Pound float64

func main() {
	var test Kilogram = 20
	testLb := Pound(test)
	fmt.Printf("Pound is %v and kilogram is %v", test, testLb)
}
```

Answer: `Pound is 20 and kilogram is 20`. It is done based only on float64

### 7. Variables and files initialization order in golang 

Files are initialized in sorted order. Package variables are initialized from top to bottom, but
whe variable has dependency on another - then it another variable will be initialized first.

```go
package main

var a = b + c // a initialized third, to 3
var b = f() // b initialized second, to 2, by calling f
var c = 1 // c initialized first, to 1

func f() int { return c + 1 }
```

There might be multiple init functions `func init() { /* ... */ }` that will be executed on startup

### 8. Package initialization

One package is initialized at a time, in order of the import of the program. If package depends on
another package then that package initializes first and process goes from top to bottom

### 9. What is declaration shadowing?

When compiler encounters reference to a name it looks for this a declaration in the innermost scope 
and goes up to the universal block.

### 10. Will the package-scope variable `cwd` initialized?

```go
package main

import (
	"log"
	"os"
)

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd() failed: %v", err)
	}
	log.Printf("Woring dir is %v", cwd)
}
```

Answer: No, because function scope variable will be created

### 11. What are all types in go?
- Basic types
- Aggregate types 
- Reference types 
- Interface types