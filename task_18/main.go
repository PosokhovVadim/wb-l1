package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu  sync.Mutex
	sum int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sum++
}

type CounterAtomic struct {
	sum int64
}

func (c *CounterAtomic) Increment() {
	atomic.AddInt64(&c.sum, 1)
}

func main() {

	var wg sync.WaitGroup
	var k int
	fmt.Scanf("%d", &k)

	switch k {
	case 1:
		var c Counter
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				c.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Println(c.sum)

	case 2:
		var c CounterAtomic
		for i := 0; i < 100000; i++ {
			wg.Add(1)
			go func() {
				c.Increment()
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Println(c.sum)
	}
}
