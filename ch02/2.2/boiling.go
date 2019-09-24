package main

import "fmt"

const boilingPoint = 212.0

func main() {
	var f = boilingPoint
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
