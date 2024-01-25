# Reflection

### Define reflection?
Reflection is just a mechanism to examine the type and value pair stored inside an interface variable

### Name and explain 1 law of reflection?
Reflection goes from interface value to reflection object.
we can get reflection type like:
`
// TypeOf returns the reflection Type of the value in the interface{}.
func TypeOf(i interface{}) Type
func ValueOf(i any) Value
`
and as we see, we can pass any value to get reflection type

### Name and explain 2 law of reflection?
Reflection goes from reflection object to interface value.
In short, the Interface method is the inverse of the ValueOf function, except that its result is always of static type interface{}.
```go
func main() {
    v := reflect.ValueOf(9.0)
    y := v.Interface().(float64) // y will have type float64.
    fmt.Println(y) //9
}
```

### Name and explain 3 law of reflection?
To modify a reflection object, the value must be settable.
Object is settable if we pass pointer of the value to reflection 
`
p := reflect.ValueOf(&x)
v := p.Elem()
fmt.Println("settability of v:", v.CanSet()) //true
`
