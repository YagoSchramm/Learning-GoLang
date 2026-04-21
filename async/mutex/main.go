package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter uint64
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			for i := 0; i < 1000; i++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Print(counter)
}
