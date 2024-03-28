package main

import (
	"sync"
)

func doRequest(v int) {
	// do request to some external API over network
}

func main() {
	ids := []int{1, 3, 5, 8, 13, 21}
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 3)
	for _, v := range ids {
		wg.Add(1)
		ch <- struct{}{}
		go func(id int) {
			defer func() {
				<-ch
			}()
			doRequest(id)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
