package main

import "fmt"

func main() {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arrCh := make(chan int64)
	resCh := make(chan int64)

	go func() {
		defer close(arrCh)
		for _, num := range arr {
			arrCh <- num
		}
	}()

	go func() {
		defer close(resCh)
		for num := range arrCh {
			result := num * 2
			resCh <- result
		}
	}()

	for result := range resCh {
		fmt.Println(result)
	}

}
