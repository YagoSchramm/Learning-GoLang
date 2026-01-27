package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	var p1 Pessoa = Pessoa{Nome: "Yago", Idade: 17}
	fmt.Println(p1)
}
