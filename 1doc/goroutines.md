# Goroutines

### CSP stands for?
communicating sequential processes

### How to make a chanel?
`ch := make(chan int)`

### Are channels comparable? 
Channels are same with == operator if they point to same underlying data structure or if they are both nil 

### What will print this?
```go
package main

import "fmt"

func channelComp() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println(ch1 == ch2)
}
```
False, even though channels have same content

### Operations with channels?
```go
package main
import "fmt"

func channelUsage()  {
	ch1 := make(chan int, 3) //create with buffer 3
	ch1 <- 1 //add
	ch1 <- 2
	x := <- ch1 //receive to var
	<- ch1 //discard
	close(ch1) //close channel 
	fmt.Println(x) 
}
```

### Describe close operation?
Channel can be closed, in that case adding to it will cause runtime panic. 
But pollers can poll from it, if elements are present there they will receive them or nil value for type otherwise.


### How to check if channel is closed?

There are 2 possible implementation to iterate through the chanel until it is closed:
1. use 2 return values when retrieving from channel
```go
for {
	x, ok := <- naturals 
	if !ok {
	    break // channel was closed and drained	
    }
	squares <= x * x // do work otherwise 
}
```
2. Iterate thought a channel and for look will check it
```go
for x := range naturals {
    square <- x * x
}
```

### When channel closing causes panic? 
Closing already closed channel or closing nil channel

### Describe unidirectional send-only channel types
This is a channel type when only sending and closing are 
permitted. Receiving and iterating will cause compile error
Example:
`out chan<- int`

### Describe unidirectional receive-only channel types
Only receiving from channel allowed, closing the channel will also cause compile error.
Example
`in <- chan int`

### Are unbuffered channels blocking?
Yes, unbuffered chanel(created like `make(chan int)`) is also called 
synchronous, and it blocks if:
- Trying to receive from empty channel until add operation from other goroutine
- Sending to the channel blocks until client receives from it 

### Why this code snippet will have a deadlock?
Assume function receives channel filenames, 
that is closed, but we don't know number of entries. thumbnail function 
should run in parallel, and we want to wait for all thumbnails to 
finish before computing total sizes number.

```go
func makeThumbnails(filenames <-chan string) {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			size := thumbnail(filename)
			fmt.Printf("Generated file for thumbnail %s \n", filename)
			sizes <- size
		}(f)
	}
	wg.Wait()
	close(sizes)

	var total int64
	for size := range sizes {
		fmt.Printf("Adding size %s \n", size)
		total += size
	}
}
```
This code will not run because we create sizes channel, and it is
unbuffered channel. That means after second iteration of loop over
filenames code execution will stop waiting for sizes channel to be empty, but that 
will never happen because wg.Wait() will stop in another thread and prevent
execution of loop over sizes slice.

### Explain this select statement?
```go
select {
    case <-tick:
        fmt.Println(countdown)
    case <-abort:
        fmt.Println("Aborting")
		return
}
```

Select statement with channels executes in following way:
- blocks if no goroutines can execute (adding to the full, querying from empty). In this case it will be blocked if both are empty
- executes if only 1 is ready, in our case it will execute tick or about depending on which one has entry
- if both are ready - picks randomly. in our case almost imposible 

### How can we canncel goroutine? 
For this we need function that will return true to every goroutine if it was cancelled
In order to achieve this we need to create a channel and then check if it was closed. If it was closed -> then cancelled, 
otherwise not:
```go
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
```
Then we just need to close channel done when needed


### Datarace definition? 
A data race occurs whenever two goroutines access the same variable concurrently and at least one of the accesses is a write.

### Is sync.Mutex in golang reentrant?
No, it is not

### Descrive sync.RWMutex?
It is read write lock that can be described as multiple readers, single writer lock

### Why do we need sync.Once?
This object can be used for initialization and verifying that action has taken place exactly once and finished before
second invocation. 
example

```go
package main

import (
	"image"
	"sync"
)
var initialize sync.Once

func getIcon(filename string) image.Image {
	initialize.Do(loadIcons)
	return icons[filename]
}
```

## What will do this code?
```go
func main() {
	ch := make(chan struct{})
	go func() {
		_, ok := <-ch
		fmt.Println(ok)
	}()
	time.Sleep(1 * time.Second)
}
```
Answer: it will print nothing because created goroutine will be blocked when trying to read from channel

