package main

import "time"

func fun(val string) {
	for i := 0; i < 3; i++ {
		println(val)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go fun("a1")
	go fun("a2")
	go fun("a3")
	time.Sleep(10 * time.Second)

}
