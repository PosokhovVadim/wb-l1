package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ch := make(chan int)

	//read data
	go func(ch <-chan int) {
		for value := range ch {
			fmt.Println(value)
		}
	}(ch)

	//send data
	go func(ch chan<- int, n int) {
		for i := 0; i < n; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}(ch, n)

	time.Sleep(time.Duration(n) * time.Second)

	close(ch)
}
