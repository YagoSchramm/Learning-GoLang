package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const N = 10
	var values [N]string
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d)
			cond.L.Lock()
			values[i] = string('a' + i)
			cond.L.Unlock()
			cond.Signal()
		}(i)
	}
	var checkCondition = func() bool {
		fmt.Println(values)
		for _, value := range values {
			if value == "" {
				return false
			}
		}
		return true
	}
	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait()
	}
}
