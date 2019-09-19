package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var s, sep string
	for c := 0; c < 100000; c++ {
		s = ""
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}

	}

	fmt.Println(s)
	fmt.Println("elapsed time:", time.Since(start))

}
