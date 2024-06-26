package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 3, 8, 7, 11, 2, 15}

	sort.Ints(arr)

	fmt.Println(arr)
}
