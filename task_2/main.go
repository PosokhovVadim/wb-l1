package main

import (
	"fmt"
	"sync"
	"time"
)

func first_solution(arr []int64) {
	var wg sync.WaitGroup

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(arr[i] * arr[i])
		}(i)
	}

	wg.Wait()
}

func second_solution(arr []int64) {

	for i := 0; i < len(arr); i++ {
		go func(i int) {
			fmt.Println(arr[i] * arr[i])
		}(i)
	}

	time.Sleep(time.Microsecond * 75)
}

func main() {
	arr := []int64{2, 4, 6, 8, 10}

	fmt.Println("First solution: ")
	first_solution(arr)

	fmt.Println("Second solution: ")
	second_solution(arr)
}
