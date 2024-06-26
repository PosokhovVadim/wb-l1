package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("working, time: ", time.Now().Second())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func contextMethod() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	fmt.Println("all channels closed")

}

func stopSignalMethod() {
	ch := make(chan bool)

	go func(ch <-chan bool) {
		for {
			select {
			case <-ch:
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("working, time: ", time.Now().Second())
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ch)

	time.Sleep(2 * time.Second)
	ch <- true
	fmt.Println("channel closed")

}

func closeChanMethod() {
	ch := make(chan int)

	go func(ch <-chan int) {

		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("channel closed")
					return
				}
			default:
				fmt.Println("Worker is working...")
				time.Sleep(500 * time.Millisecond)
			}

		}

	}(ch)

	time.Sleep(2 * time.Second)

	close(ch)

	fmt.Println("channel closed")
}

func main() {
	var k int
	fmt.Scanf("%d", &k)

	switch k {
	case 1:
		contextMethod()
	case 2:
		stopSignalMethod()
	case 3:
		closeChanMethod()

	default:
		closeChanMethod()
	}
}
