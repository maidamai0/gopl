package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	conuts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		conuts[input.Text()]++
	}

	for line, count := range conuts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}

}
