package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{6, 3, 8, 7, 11, 2, 15}
	sort.Ints(arr)

	target := 6

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found\n", target)
	}
}
