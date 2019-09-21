package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, 0, counts)
	} else {
		for index, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}

			countLines(f, index, counts)
			f.Close()

		} // read file
	} // iterate files

	// print
	for line, indeies := range counts {
		if len(indeies) > 1 {
			fmt.Printf("%s\t:", line)

			for _, index := range indeies {
				fmt.Printf("%s ", files[index])
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, index int, counts map[string][]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], index)
	}
}
