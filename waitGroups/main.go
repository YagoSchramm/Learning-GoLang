package main

import (
	"fmt"
	"sync"
	"time"
)

func DemoraT(t time.Duration) {
	fmt.Println("INICIO DA EXECUCAO")
	time.Sleep(t)
	fmt.Printf("FIM DA EXECUCAO DEPOIS DE %s \n", t)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DemoraT(time.Duration(i * int(time.Second)))
		}()

	}
	wg.Wait()
}
