package main

import "fmt"

func main() {
	x := 2
	y := 5

	x, y = y, x

	fmt.Println(x, y)
}
