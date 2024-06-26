package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int64{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	s := int64(0)
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s = s + arr[i]*arr[i]
		}(i)
	}

	wg.Wait()
	fmt.Printf("sum = %d", s)
}
