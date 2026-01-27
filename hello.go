package main

import "fmt"

var pontuacao int
var pontuacao_quebrada float64
var jogador string
var is_alive bool
var globalVar float64

func main() {
	var localVar float64 = 9.4
	localVar2 := 9.2
	jogador = "Yago"
	pontuacao = 10
	pontuacao_quebrada = 9.5
	is_alive = true
	fmt.Println("Hello go world!")
	fmt.Println(jogador)
	fmt.Println(pontuacao)
	fmt.Println(pontuacao_quebrada)
	fmt.Println(is_alive)
	fmt.Println(localVar)
	fmt.Println(localVar2)
	fmt.Println("Insira o primeiro número:")
	var numero1 float64
	fmt.Scanln(&numero1)
	fmt.Println("Insira o segundo número:")
	var numero2 float64
	fmt.Scanln(&numero2)
	soma := numero1 + numero2
	fmt.Printf("A soma dos dois números é: %f\n", soma)
}
