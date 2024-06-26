package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, workerID int, wg *sync.WaitGroup, mainCh <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case num := <-mainCh:
			fmt.Printf("Worker %d received %d\n", workerID, num)
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	var wg sync.WaitGroup

	mainCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg, mainCh)
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stopCh
		fmt.Println("\nReceived an interrupt, stopping workers...")
		cancel()
	}()

	go func() {
		for {
			r := rand.Intn(100)
			mainCh <- r
			time.Sleep(500 * time.Millisecond)

		}
	}()
	wg.Wait()

}
