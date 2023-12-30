package goroutines

import "fmt"

func Pipeline() {
	naturals, squares := make(chan int), make(chan int)
	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}

func counter(out chan<- int) { //send only channel
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
