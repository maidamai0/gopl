package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	rsp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	len, err := io.Copy(ioutil.Discard, rsp.Body)
	rsp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("reading %s failed:%s", url, err)
		return
	}

	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, len, url)
}
