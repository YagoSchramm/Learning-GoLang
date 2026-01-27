package main

import "fmt"

func soma(x *int) {
	*x = *x + 2
}
func main() {
	fmt.Println("Digite o primeiro número")
	var n1 int
	fmt.Scan(&n1)
	soma(&n1)
	fmt.Println(n1)
}
