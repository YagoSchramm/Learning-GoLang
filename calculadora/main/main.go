package main

import calculadora "aprendendoGoLang/calculadora/structs"

func main() {
	calc := &calculadora.Calculadora{}
	runner := calculadora.Runner{Calculadora: calc}

	runner.Execute()

}
