package main

import (
	"fmt"
	"sync"
)

// for case 1
func writeToSyncMap(wg *sync.WaitGroup, data *sync.Map, key int, value int) {
	defer wg.Done()
	data.Store(key, value)

}

// for case 2
type CMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func NewCMap() *CMap {
	return &CMap{
		m: make(map[int]int),
	}
}

func (c *CMap) Add(key int, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

func (c *CMap) PrintMap() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for k, v := range c.m {
		fmt.Println(k, v)
	}

}

func main() {
	var k int
	fmt.Scanf("%d", &k)

	switch k {
	case 1:
		var data sync.Map
		var wg sync.WaitGroup
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go writeToSyncMap(&wg, &data, i, i)
		}
		wg.Wait()
		data.Range(func(key, value interface{}) bool {
			fmt.Println(key, value)
			return true
		})

	case 2:
		c := NewCMap()
		var wg sync.WaitGroup
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				c.Add(i, i)
			}(i)

		}
		wg.Wait()
		c.PrintMap()
	}

}
