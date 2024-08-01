package main

/*
Os ponteiros são variáveis especiais que armazenam o endereço de memória de outra variável, em vez de armazenar um valor diretamente.
Você pode pensar neles como "indicadores" que apontam para onde outra variável está armazenada na memória do computador.

Os ponteiros são úteis por várias razões:

1. Permitem modificar variáveis de forma indireta.
2. Podem tornar os programas mais eficientes, especialmente ao trabalhar com grandes quantidades de dados.
3. São necessários para criar estruturas de dados complexas como listas encadeadas e árvores.

Em Go, usa-se o símbolo & para obter o endereço de memória de uma variável, e o símbolo * para declarar um ponteiro ou para acessar o valor para o qual ele aponta.
*/

// func main() {
// 	// Declaramos una variable normal
// 	edad := 30

// 	// Creamos un puntero a la variable edad
// 	punteroEdad := &edad

// 	// Imprimimos el valor y la dirección de memoria
// 	fmt.Println("Valor de edad:", edad)
// 	fmt.Println("Dirección de memoria de edad:", punteroEdad)

// 	// Modificamos el valor a través del puntero
// 	*punteroEdad = 31

// 	// Imprimimos el nuevo valor
// 	fmt.Println("Nuevo valor de edad:", edad)
// }

func modificarIdade(idade *int) {
	*idade = *idade + 1
}

// func main() {
// 	// Declaramos uma variável
// 	idade := 30

// 	fmt.Println("Idade antes:", idade)

// 	// Chamamos a função passando o endereço da variável
// 	modificarIdade(&idade)

// 	fmt.Println("Idade depois:", idade)
// }
