package main

import (
	"fmt"
	"time"
)

func getValueFromChannel(ch chan<- string, t time.Duration, i int) {
	time.Sleep(t)
	ch <- "Done this channel" + fmt.Sprintln(i)
}
func main() {
	ch := make(chan string)
	go getValueFromChannel(ch, 2*time.Second, 1)
	ch2 := make(chan string)
	go getValueFromChannel(ch2, 2*time.Second, 2)
	ch3 := make(chan string)
	go getValueFromChannel(ch3, 2*time.Second, 3)
	select {
	case rtval := <-ch:
		fmt.Println(rtval)
		break
	case rtval := <-ch2:
		fmt.Println(rtval)
		break
	case rtval := <-ch3:
		fmt.Println(rtval)
		break
	}
}
