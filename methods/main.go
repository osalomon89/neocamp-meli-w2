package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saludar() {
	fmt.Printf("Ola, mi nombre es %s y tengo %d años.\n", p.Nome, p.Idade)
}

func (p *Pessoa) CumplirAnos() {
	p.Idade += 1
}

func main() {
	// Crear una nueva persona
	persona := Pessoa{Nome: "Juan", Idade: 30}

	// Llamar al método Saludar
	persona.Saludar() // Output: Hola, mi nombre es Juan y tengo 30 años.

	// Llamar al método CumplirAños
	persona.CumplirAnos()

	// Verificar la edad actualizada
	persona.Saludar() // Output: Hola, mi nombre es Juan y tengo 31 años.
}
