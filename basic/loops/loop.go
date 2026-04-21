package main

import "fmt"

func main() {
	n := 0
	for i := 1; i <= 10; i++ {
		fmt.Println("Número:", i)
		fmt.Scan(&n)
		if n%2 == 0 {
			fmt.Println("o número é par")
		} else {
			fmt.Println("o número é impar")
		}
	}
}
