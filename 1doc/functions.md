# Functions 

### 1. Why functions in go allow unnamed parameter?
Unnamed parameters in function may be useful to declare, implement some interface. Unnamed parameters show that they are not used.

### 2. Is function overloading possible in go? 
No, this functions will not compile:
```go
func add(x int, y int) int {
	
}

func add(x float64, y float64) float64 {
	
}
```

### 3. What is the size of stack in Go?
Go stack has unlimited size, so recursion can grow without stackOverflow

### 4. Can we straight return multivalued function result or pass this result to another function? 
Yes if all multivalued parameters have same type

### 5. What are the bare returns?
return the named vars that are initialized while function invoked
```go
func bareReturn() (x, y int) { //initializes x, y to default values in function invocation
	if rand.Int() < 100 {
		return //return 0, 0
	}
	x = rand.Int()
	y = rand.Int()
	return //returns x, y
}
```

### 6. Why go function return an error instead of `throw`ing an exception?
Because exception often contains big stack trace which cannot describe error fully. Error demands more effort in handling 
but provides more descriptive and short messages. 

### 7. Common error handling strategies?
- Propagate same error:
```go
func fetchUrl(url string) (int, error) {
    res, err := http.Get(url)
    if err != nil {
        return 0, err //return same error
    }
    return res.StatusCode, nil
}
```
- Format to create new error with additional parameters
```go
func fetchUrl(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("Fetch  failed for url %s, error %v", url, err) //creating new error with some info
	}
	return res.StatusCode, nil
}
```
- Implement retry mechanism (maybe with exponential backoff)
- Print error and stop program gracefully, that should be done only in main package. For this `log.Fatalf()` or `os.Exit(1)` can be used
- Log error and proceed possibly with reduced functionality
- Ignore error in rare cases

### 8. Can we assign function to the variable?
Yes, because functions are first-class values

### 9. What will print this code?
```go
func fetchUrl(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("Fetch  failed for url %s, error %v", url, err)
	}
	return res.StatusCode, nil
}

func functionValues() {
	f := fetchUrl
	fmt.Printf("%T", f)
}
```
Answer: `func(string) (int, error)`

### 10. What belongs to the function type?
Order and types of input and output parameters

### 11. What will print this code?
```go
package main

import "fmt"

func main() {
	squareInstance := squares()
	fmt.Println(squareInstance()) 
	fmt.Println(squareInstance())
	fmt.Println(squareInstance()) 
}

func squares() func() int {
	x := 0
	return func() int {
		x++
		return x * x
	}
}
```

Answer: It will print 1, 4, 9
Every instance of anonymous function returned by square() will have own instance of variable x

### 12. Provide minimum example for anonymous function
```go
package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println("message")
	}("test")
}
```

### 13. What does this code print and why?
```go
func ShowFunctionClousure() {
	strs := []string{"str1", "str2", "str3"}
	var prints []func() string
	for _, str := range strs {
		prints = append(prints, func() string {
			return str
		})
	}
	fmt.Println(prints[0]()) //what does it print
}
```

This code prints `str3`, because `str` is a variable managed by a for loop. 
This means that anonymous functions captures this variable and this variable address is changed
while iteration is done. In order to allocate individual value for each function new variable should be defined:
```go
func ShowFunctionClousure() {
	strs := []string{"str1", "str2", "str3"}
	var prints []func() string
	for _, str := range strs {
		strCurr := str
		prints = append(prints, func() string {
			return strCurr
		})
	}
	fmt.Println(prints[0]()) //prints str1
}
```


### 14. how variadic function argument works under the hood?
It implicitly makes caller to allocate an array of arguments, copy all arguments in this array
and then pass a slice of entire array to the function
These 2 are same
```go
package main

import (
	"fmt"
	"strings"
)

func Variadic() {
	concat("a", "b", "c")
	//same as
	values := []string{"a", "b", "c"}
	concat(values...)
}

func concat(stringsArgs ...string) string {
	fmt.Printf("Type of string is %T", stringsArgs) //Type of string is []string
	return strings.Join(stringsArgs, " ")
}

```

### 15. When the expressions in defer statement are evaluated? When statement is executed?
Expressions are evaluated instantly, but statements are executed right after the end of execution of a function.

### 16. What if multiple defer statements are present? 
They are executed in reverse order from which they were deferred 

### 17. What will print this code?
```go
package functions

import "fmt"

func TestDefer() {
	a := 1
	defer fmt.Printf("defer 1, value %v \n", a)
	a = 1
	defer fmt.Printf("defer 2, value %v \n", a)
	sideFunction(3)
	a = 4
	defer fmt.Printf("defer 4, value %v \n", a)
	panic("acd")
}

func sideFunction(a int) error {
	defer fmt.Printf("defer in method, value %v \n", a)
	err := fmt.Errorf("error in fucntion %v \n", "sideFunction")
	return err
}
```

Answer:
defer in method, value 3
defer 4, value 4
defer 2, value 1
defer 1, value 1
panic: acd

### 18. Defer is async or sync?
sync:
```go
package main

import (
	"fmt"
	"time"
)

func TestDefer() {
	deferWithSleep()
	fmt.Println("After deferWithSleep")
}

func deferWithSleep() {
	defer func() {
		time.Sleep(time.Second * 2)
		fmt.Println("defer execution")
	}()
	fmt.Println("deferWithSleep")
}
```
This will print
deferWithSleep
defer execution
After deferWithSleep

### 19. Can defer change the result of a function?
Yes, if result is a named variable
````go
package main

func TestDefer() {
	deferChangeResult()// will return 10
}

func deferChangeResult() (a int) {
    defer func() { a *= 10 }()
    return 1
}

````

### 20. What are the best scenarios to use panic?
Only in situation that causes bugs and not expected behaviour (like index out of bound)


### 21. What Must prefix for the function means?
That function should execute without exception. If this is not the case - panic.
Use case - mustCompile regex, that is known beforehand

### 22. Will be deferred executed after panic?
yes
```go
func deferChangeResult() (a int) {
	defer func() { fmt.Println("After panic") }()
	panic("Test")
	return 1
}
```
This will print `After panic`
and then stacktrace

### 23. What this function will print?
```go
package main

import "fmt"

func TestDefer() {
	fmt.Println(testPanic())
}

func testPanic() (a int) {
	defer func() {
		if p := recover(); p != nil {
			a = 1
		}
	}()
	panic("ABCD")
}
```

    Answer: It prints `1`, no panic
