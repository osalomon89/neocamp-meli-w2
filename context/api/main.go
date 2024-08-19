package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fazerRequisicao(ctx context.Context) (*http.Response, error) {
	// Cria uma nova requisição HTTP
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
	if err != nil {
		return nil, err
	}

	// Envia a requisição e retorna a resposta
	client := &http.Client{}
	return client.Do(req)
}

func main() {
	// Define um timeout de 2 segundos para a requisição
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Garante que o contexto seja cancelado após o uso

	// Faz a requisição à API
	resp, err := fazerRequisicao(ctx) // Essa URL simula um atraso de 3 segundos
	if err != nil {
		// Verifica se o erro foi devido ao contexto ter sido cancelado
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Requisição cancelada: Tempo limite excedido")
		} else {
			fmt.Println("Erro na requisição:", err)
		}
		return
	}
	defer resp.Body.Close()

	// Se a requisição for bem-sucedida, exibe o status
	fmt.Println("Requisição bem-sucedida:", resp.Status)
}
