package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Digite sua idade")
	var idade int
	fmt.Scan(&idade)
	var anoAtual int = time.Now().Year()
	anoNascimento := anoAtual - idade
	fmt.Printf("Você nasceu em: %d\n", anoNascimento)
}
