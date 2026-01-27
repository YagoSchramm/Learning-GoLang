package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["Yago"] = 17
	for key, value := range map1 {
		fmt.Println("Chave", key, "Valor:", value)
	}
}
