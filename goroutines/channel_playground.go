package goroutines

import "fmt"

func ChannelEntrypoint() {
}

func channelComp() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println(ch1 == ch2)
}

func channelUsage() {
	ch1 := make(chan int)
	ch1 <- 1
	ch1 <- 2
	x := <-ch1
	<-ch1
	close(ch1)
	fmt.Println(x)
}

func intGen() int {
	return 1
}
