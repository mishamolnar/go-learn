package goroutines

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CountDown() {
	abort := make(chan struct{})
	go func() { //terminate on any input
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			log.Fatal(err)
		}
		abort <- struct{}{}
	}()
	fmt.Println("Starting countdown, type enter to abort")
	tick := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown >= 0; countdown-- {
		select {
		case <-tick.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Aborting")
			return
		}
	}
	tick.Stop()
	fmt.Println("Launch")
}
