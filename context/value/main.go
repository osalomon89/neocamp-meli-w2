package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Definindo uma chave para o contexto
const requestIDKey = "reqID"

// Middleware para adicionar o ID da requisição ao contexto
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Gerar um ID único para a requisição (aqui usamos um exemplo simples)
		reqID := "12345"

		// Adiciona o ID ao contexto da requisição
		ctx := context.WithValue(r.Context(), requestIDKey, reqID)

		// Passa a requisição para o próximo handler, com o novo contexto
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler que processa a requisição e acessa o ID da requisição a partir do contexto
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Recupera o ID da requisição do contexto
	reqID := r.Context().Value(requestIDKey).(string)

	// Usa o ID da requisição para log ou resposta
	log.Printf("Processando requisição com ID: %s", reqID)
	fmt.Fprintf(w, "Olá, mundo! Requisição ID: %s\n", reqID)
}

func longRunningProcess(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Iniciando processo longo...")

	select {
	case <-time.After(5 * time.Second): // Simula uma operação que leva 10 segundos
		fmt.Fprintln(w, "Processo concluído")
	case <-ctx.Done(): // O contexto foi cancelado
		log.Println("Processo cancelado pelo cliente")
		http.Error(w, "Requisição cancelada pelo cliente", http.StatusRequestTimeout)
		return
	}

	log.Println("Processo longo finalizado")
}

func main() {
	// Cria um novo mux (roteador)
	mux := http.NewServeMux()

	// Adiciona o middleware e o handler ao roteador
	mux.Handle("/", requestIDMiddleware(http.HandlerFunc(helloHandler)))

	mux.Handle("/process", http.HandlerFunc(longRunningProcess))

	// Inicia o servidor HTTP
	log.Println("Servidor rodando na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
