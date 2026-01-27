package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Por favor, insira sua idade:")
	fmt.Scanln(&idade)
	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	} else {
		fmt.Println("Você é maior de idade.")
	}
}
