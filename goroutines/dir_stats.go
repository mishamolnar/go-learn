package goroutines

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DirWalker(dirs []string) {
	tick := time.NewTicker(250 * time.Millisecond)
	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	go func() { //interrupt function
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	for i := range dirs {
		wg.Add(1)
		go func(dir string) {
			walkDir(dir, &wg, fileSizes)
		}(dirs[i])
	}
	var totalBytes, totalFiles int64
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

loopOverSizes:
	for {
		select {
		case <-done:
			for _ = range fileSizes {
				//do nothing, discard all
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loopOverSizes
			}
			totalBytes += size
			totalFiles++
		case <-tick.C:
			fmt.Printf("%d files  %.1f MB\n", totalFiles, float64(totalBytes)/1e6)
		}
	}

	fmt.Printf("Total: %d files  %.1f MB\n", totalFiles, float64(totalBytes)/1e6)
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirEntries(dir) {
		time.Sleep(10 * time.Millisecond)
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, wg, fileSizes)
		} else {
			fileInfo, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting an error for file %s \n", entry.Name())
			}
			fileSizes <- fileInfo.Size()
		}
	}
}

var semaphore = make(chan struct{}, 20)

func dirEntries(dir string) []os.DirEntry {
	semaphore <- struct{}{}
	defer func() { <-semaphore }() //semaphore wrapper that allows only 20 entries at the same time
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v \n", err)
		return entries
	}
	return entries
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
