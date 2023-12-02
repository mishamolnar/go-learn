# Questions basic data types

### 1. What are the basic data types?
- numbers, strings and booleans

### 2. What is the size of uint and int?
32 or 64 bits depending on the compiler

### 3. All arithmetic, logic and comparison operators?
This is in order of decreasing precedence 
- `* / % << >> & &^`
- `+ - | ^`
- `== != < <= > >=`
- `&& ||`

### 4. Are basic data types comparable and ordered using standard operators?
Yes

### 5. Explain all bitwise operations
- `<<` - left shift 
- `>>` - right shift
- `&` - bitwise AND
- `|`  -bitwise OR
- `&^` - bit clear AND NOT
- `^` - bitwise XOR

### 6. NaN. When occurs and how to test?
- NaN when 0 / 0 or sqrt(-1)
- NaN should be tested with `math.IsNaN`
- NaN returns false in any comparison, even with itself `fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"`

### 7. Are strings immutable? 
yes, strings are immutable in go. Operation like `s[0] = 'a'` is not allowed
This makes possible to share the memory for 2 different variables

### 8. What should be used to append strings?
bytes.Buffer

### 9. What can be types of constants?
boolean, string or number

### 10. What is the constant pros compared to standard vars?
- Greater numeric precision, at least 256 bits
- They can be used in more context with explicit casting

```go
package main

import (
	"fmt"
	"math"
)

const (
	e  = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}
```

### 11. What is iota and how can be used?

```go
package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(YiB / ZiB) // no overflow, even thought ZiB > 64 bits number. Output: "1024"
}
```
