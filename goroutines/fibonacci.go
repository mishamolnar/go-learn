package goroutines

import (
	"fmt"
	"time"
)

func CalculateFibonacci(num int) {
	go spinner(100 * time.Millisecond)
	fibN := fib(num)
	fmt.Printf("\r Fibonacci(%d) = %d\n", num, fibN)
}

func spinner(delay time.Duration) {
	loading := []byte("..............")
	for {
		for i := range loading {
			loading[i] = '|'
			fmt.Printf("\r%s", string(loading))
			time.Sleep(delay)
		}
		loading = []byte("..............")
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
