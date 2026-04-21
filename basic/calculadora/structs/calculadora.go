package calculadora

import (
	"errors"
)

type Calculadora struct {
	n1 float32
	n2 float32
}

func (calc Calculadora) Somar() float32 {
	return calc.n1 + calc.n2
}
func (calc Calculadora) Diminuir() float32 {
	return calc.n1 - calc.n2
}

func (calc Calculadora) Dividir() (float32, error) {
	if calc.n2 == 0 {
		return 0, errors.New("divisão de por zero não é permitida")
	}
	return calc.n1 / calc.n2, nil

}

func (calc Calculadora) Multiplicar() float32 {
	return calc.n1 * calc.n2
}
