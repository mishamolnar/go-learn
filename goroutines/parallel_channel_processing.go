package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
)

func RunThumbnails() {
	thumbnails := make(chan string, 10)
	generateFiles(thumbnails)
	makeThumbnails(thumbnails)
}

func generateFiles(ch chan<- string) {
	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("file %d", i)
		ch <- str
		fmt.Println("Submitted file ", str)
	}
	close(ch)
}

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

func thumbnail(filename string) int64 {
	return rand.Int63n(5000)
}
