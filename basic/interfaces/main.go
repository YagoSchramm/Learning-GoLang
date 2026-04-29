package main

import "fmt"

type CalculadoraInterface interface {
	Somar(n1, n2 int) int
	Diminuir(n1, n2 int) int
	Dividir(n1, n2 int) int
	Multiplicar(n1, n2 int) int
}

func NewCalculadoraInterface() CalculadoraInterface {
	return &Calculadora{}
}

type Calculadora struct {
}

func (c *Calculadora) Somar(n1, n2 int) int {
	return n1 + n2
}
func (c *Calculadora) Diminuir(n1, n2 int) int {
	return n1 - n2
}
func (c *Calculadora) Dividir(n1, n2 int) int {
	return n1 / n2
}
func (c *Calculadora) Multiplicar(n1, n2 int) int {
	return n1 * n2
}
func main() {
	calc := NewCalculadoraInterface()
	fmt.Print(calc.Somar(1, 2))

}
