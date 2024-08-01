package main

// Definimos uma struct Pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Função que recebe um ponteiro para Pessoa e modifica seus dados
func fazerAniversario(p *Pessoa) {
	p.Idade++
	p.Nome = p.Nome + " Neymar"
}

// func main() {
// 	// Criamos uma instância de Pessoa
// 	pessoa := Pessoa{
// 		Nome:  "João",
// 		Idade: 30,
// 	}

// 	fmt.Println("Antes:", pessoa)

// 	// Chamamos a função passando o endereço da struct
// 	fazerAniversario(&pessoa)

// 	fmt.Println("Depois:", pessoa)
// }
