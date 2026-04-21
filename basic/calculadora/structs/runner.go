package calculadora

import (
	"errors"
	"fmt"
)

type Runner struct {
	Calculadora *Calculadora
	operacao    string
}

func (r *Runner) SolicitarValores() error {
	fmt.Println("Digite o primeiro número:")
	if _, err := fmt.Scan(&r.Calculadora.n1); err != nil {
		return errors.New("Entrada do primeiro número inválida")
	}
	fmt.Println("Digite o segundo número:")
	if _, err := fmt.Scan(&r.Calculadora.n2); err != nil {
		return errors.New("Entrada do primeiro número inválida")
	}
	return nil
}
func (r *Runner) SolicitarOperacao() error {
	fmt.Println("Digite a operação (+,-,*,/):")
	var operacao string
	if _, err := fmt.Scan(&operacao); err != nil {
		return errors.New("Entrada da operação inválida")
	}

	switch operacao {
	case "+", "-", "*", "/":
		r.operacao = operacao
		return nil
	default:
		return errors.New("Entrada de operação inválida")

	}
}
func (r *Runner) ExecutarOperacao() {
	switch r.operacao {
	case "+":
		fmt.Println("Resultado:", r.Calculadora.Somar())
	case "-":
		fmt.Println("Resultado:", r.Calculadora.Diminuir())
	case "/":
		resultado, err := r.Calculadora.Dividir()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Resultado:", resultado)
		}

	case "*":
		fmt.Println("Resultado:", r.Calculadora.Multiplicar())
	}
}
func (r *Runner) Execute() {
	for {
		if err := r.SolicitarValores(); err != nil {
			fmt.Println(err)
			continue
		}
		if err := r.SolicitarOperacao(); err != nil {
			fmt.Println(err)
			continue
		}
		r.ExecutarOperacao()
	}
}
func NewRunner(c *Calculadora) *Runner {
	return &Runner{Calculadora: c}
}
