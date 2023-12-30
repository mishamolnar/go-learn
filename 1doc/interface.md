# Interface

### 1. Is interface a concrete type?
No, it's interface type

### 2. Interface example?
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### 3. Interface embedding example? 
```go
// ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

### 4. What is the main advantage of interfaces in go?
Satisfied implicitly
That means to satisfy and use interface - we don't need to mention this interface in type:
Here is assosiated os.Write interface with os/User
```go
package main

import (
	"fmt"
	"os/user"
)

type CountWriter user.User

func (c *CountWriter) Write(p []byte) (n int, err error) {
	c.Name += string(p)
	return len(p), nil
}

func WriterUsage() {
	var wr CountWriter
	fmt.Fprint(&wr, "Something")
	fmt.Fprint(&wr, " Something")
	fmt.Println(wr.Name)
}

```

### 5. What is interface value and what does it contain?
Interface variable contains always dynamic types and dynamic values
```go
package main

import (
	"io"
    "os"
    "bytes"
)

func main()  {
	var w io.Writer //dynamic value - null, dynamic type null
	w=os.Stdout //dynamic value - file for stdout, dynamic type *os.File
	w=new(bytes.Buffer) //dynamic value - []byte, dynamic type bytes.Buffer
	w=nil //dynamic value - null, dynamic type null
}
```

### 6. When interfaces can be compared?
Interfaces can be compared if their dynamic types are equal and if values of types are equals as well

### 7. What if interfaces are holding incomparable types?
Runtime panic
```go
package intefaces

import "fmt"

func InterVal() {
	sl := []int{1, 2, 3}
	var x interface{} = sl
	//fmt.Println(sl == sl) //compile error 
	fmt.Printf("Slices compared with result %v \n", x == x) //runtime panic
}
```

### 8. What will print this code?
```go
func testNull() {
	var buff *bytes.Buffer //buffer satisfies Writer
	f(buff)
}

func f(w io.Writer) {
	if w != nil {
		w.Write([]byte("test"))
	} else {
		fmt.Println("writer is NIL")
	}
}
```
It will fail, because `w` will be not null in function  `f`. The reason for this is that interface io.Writer will have 
dynamic value `nil`, but dynamic type will be not `nil`

### 9. Why implementation of Error interface errorString has pointer receiver?
```go
package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

Pointer receiver means that object does not implement interface, only pointer to this object implements interface.
`errorString` is not exported, which means that we will manipulate interface always
And this means that comparing 2 errors (with underlying errorString) type will always return false, because it is pointer comparison.


### 10. What is type assertions?
It is a way to cast interface value `x` to concrete type `T` or another interface `x.(T)`
- If we want to cast value to concrete type:
Then if interface value really holds this type - then it returns dynamic value (`T` value), otherwise - runtime panic
```go
package intefaces

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func interToType() {
	var w io.Writer
	w = os.Stdout           //os.Stdout is *os.File
	fl := w.(*os.File)      //ok, because w is really a file
	bf := w.(*bytes.Buffer) //panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	fmt.Println(fl, bf)
}
```
- If we want to extract another interface: then it will be checked if dynamic type satisfies new interface.
```go
package intefaces

import (
	"fmt"
	"io"
	"os"
)

func interToInterface() {
	var w io.Writer
	w = os.Stdout //os.Stdout is *os.File
	rw := w.(io.ReadWriter) //ok, because file satisfies interface
	bs := w.(io.ByteScanner) //panic
	fmt.Println(rw, bs)
}
```

### 11. Can we avoid runtime panic when doing type assertion?
Yes, for this additional output variable needed for checking
```go
package intefaces

import (
	"fmt"
	"io"
	"os"
)

func interToInterface() {
	var w io.Writer
	w = os.Stdout //os.Stdout is *os.File
	rw := w.(io.ReadWriter) //ok, because file satisfies interface
	bs, ok := w.(io.ByteScanner) //no runtime panic
	fmt.Println(ok) //false
	fmt.Println(rw, bs) //bs is <nil> because it is not a dynamic value
}
```

### 12. What will print this interface comparasion?
```go
package main

import "fmt"

func InterTesting() {
	var gr1 greet = &ComparableObject{"hello"}
	var gr2 greet = &ComparableObject{"hello"}
	fmt.Println(gr1 == gr2)
}

type ComparableObject struct {
	text string
}

type greet interface {
	Greet() string
}

func (co *ComparableObject) Greet() string {
	return co.text
}
```
This will print false, because they are pointers. It has nothing to do with interface

### 13. What is parametric polymorphism and ad-hoc polymorphism?
Parametric polymorphism - function with generic type
Ad-hoc polymorphism - function with different implementations for different types

### 14. What is type switch and where it can be used?
type switch syntax: `variable.(type)` and it returns type that can be used only in switch statement 
```go
package main

import "fmt"

func idiotFormatting(in interface{}) {
	switch x := in.(type) {
	case nil:
		fmt.Println("nil nil nil")
	case bool:
		fmt.Println(x, " is a value of boolean x")
	case int:
		b := max(x, 12)
		fmt.Println(b, " is an int")
	case string:
		fmt.Printf("This was a string %s", x)
	}
}
```