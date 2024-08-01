package main

import "fmt"

// Interface para o repositório de dados
type RepositorioProduto interface {
	BuscarProduto(id int) string
	SalvarProduto(nome string) bool
}

// Implementação fictícia para MySQL
type MySQLRepositorio struct{}

func (m *MySQLRepositorio) BuscarProduto(id int) string {
	return fmt.Sprintf("Produto %d do MySQL", id)
}

func (m *MySQLRepositorio) SalvarProduto(nome string) bool {
	fmt.Printf("Salvando '%s' no MySQL\n", nome)
	return true
}

// Implementação fictícia para MongoDB
type MongoDBRepositorio struct{}

func (m *MongoDBRepositorio) BuscarProduto(id int) string {
	return fmt.Sprintf("Produto %d do MongoDB", id)
}

func (m *MongoDBRepositorio) SalvarProduto(nome string) bool {
	fmt.Printf("Salvando '%s' no MongoDB\n", nome)
	return true
}

// Serviço que usa o repositório
type ServicoProduto struct {
	repo RepositorioProduto
}

func (s *ServicoProduto) ObterProduto(id int) string {
	return s.repo.BuscarProduto(id)
}

func (s *ServicoProduto) CriarProduto(nome string) bool {
	return s.repo.SalvarProduto(nome)
}

func main() {
	// Criando instâncias dos repositórios
	repoMySQL := &MySQLRepositorio{}
	//repoMongoDB := &MongoDBRepositorio{}

	// Criando serviços com diferentes repositórios
	servicoMySQL := &ServicoProduto{repo: repoMySQL}
	//servicoMongoDB := &ServicoProduto{repo: repoMongoDB}

	// Usando os serviços
	fmt.Println(servicoMySQL.ObterProduto(1))
	servicoMySQL.CriarProduto("Novo Produto MySQL")

	//fmt.Println(servicoMongoDB.ObterProduto(1))
	//servicoMongoDB.CriarProduto("Novo Produto MongoDB")
}
