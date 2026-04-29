package main

import calculadora "aprendendoGoLang/basic/calculadora/structs"

func main() {
	calc := &calculadora.Calculadora{}
	runner := calculadora.Runner{Calculadora: calc}

	runner.Execute()

}
