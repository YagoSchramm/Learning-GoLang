package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	on := sync.Once{}
	load := func() {
		fmt.Printf("LOADING  ......\n ")
	}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			on.Do(load)
		}()
	}
	wg.Wait()
}
