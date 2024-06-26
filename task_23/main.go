package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	indexRemove := 5

	slice = append(slice[:indexRemove], slice[indexRemove+1:]...)
	fmt.Println(slice)
}
