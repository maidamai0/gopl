package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string

	start := time.Now()
	for c := 0; c < 10000; c++ {
		s = ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}

	fmt.Println(s)
	fmt.Println("elapsed time:", time.Since(start))
}
