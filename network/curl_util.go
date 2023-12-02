package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func SendHttpParallel() {
	start := time.Now()
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchUrl(url, channel, io.Discard)
	}
	for range os.Args[1:] {
		fmt.Printf(<-channel)
	}
	elapsed := time.Since(start).Seconds()
	fmt.Printf("To fetch all urls: %f \n", elapsed)
}

func fetchUrl(url string, ch chan<- string, wrt io.Writer) {
	start := time.Now()
	resp, err := http.Get(url) //error to get
	handleError(err)
	fmt.Printf("response with code %v \n", resp.StatusCode)
	bytes, err := io.Copy(wrt, resp.Body) //err to copy
	handleError(err)
	err = resp.Body.Close()
	handleError(err) //err to close resp body
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s \n", secs, bytes, url)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error occured: %v, \n", err)
		os.Exit(1)
	}
}
