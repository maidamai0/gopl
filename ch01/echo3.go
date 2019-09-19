package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var s string
	for c := 0; c < 10000; c++ {

		s = strings.Join(os.Args[1:], " ")
	}

	fmt.Println(s)
	fmt.Println("elapsed time:", time.Since(start))
}
