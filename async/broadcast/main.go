package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	go func() {
		cond.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine 1 wait")
			cond.Wait()
		}
		fmt.Println("goroutine 1", sharedRsc)
		cond.L.Unlock()
		wg.Done()
	}()
	go func() {
		cond.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine 2 wait")
			cond.Wait()
		}
		fmt.Println("goroutine 2", sharedRsc)
		cond.L.Unlock()
		wg.Done()
	}()
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	fmt.Println("main go routine ready")
	sharedRsc = true
	cond.Broadcast()
	fmt.Println("main go routine Broadcast")
	cond.L.Unlock()
	wg.Wait()
}
