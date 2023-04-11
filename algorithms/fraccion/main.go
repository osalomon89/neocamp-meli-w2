package main

import "fmt"

type Fraccion struct {
	numerador   int
	denominador int
}

type Operable interface {
	Sumar(f Fraccion) Fraccion
	Restar(f Fraccion) Fraccion
	Multiplicar(f Fraccion) Fraccion
	Dividir(f Fraccion) Fraccion
}

func (f1 Fraccion) Sumar(f2 Fraccion) Fraccion {
	return Fraccion{
		numerador:   f1.numerador*f2.denominador + f2.numerador*f1.denominador,
		denominador: f1.denominador * f2.denominador,
	}
}

func (f1 Fraccion) Restar(f2 Fraccion) Fraccion {
	return Fraccion{
		numerador:   f1.numerador*f2.denominador - f2.numerador*f1.denominador,
		denominador: f1.denominador * f2.denominador,
	}
}

func (f1 Fraccion) Multiplicar(f2 Fraccion) Fraccion {
	return Fraccion{
		numerador:   f1.numerador * f2.numerador,
		denominador: f1.denominador * f2.denominador,
	}
}

func (f1 Fraccion) Dividir(f2 Fraccion) Fraccion {
	return Fraccion{
		numerador:   f1.numerador * f2.denominador,
		denominador: f1.denominador * f2.numerador,
	}
}

func OperarFracciones(f Fraccion, op Operable) Fraccion {
	// Realizar una operación de ejemplo. En este caso, sumar f con f.
	// Se puede cambiar la operación de acuerdo a la necesidad.
	return op.Sumar(f)
}

func main() {
	f1 := Fraccion{numerador: 1, denominador: 2}
	result := OperarFracciones(f1, f1)
	fmt.Printf("Resultado: %d/%d\n", result.numerador, result.denominador)
}
