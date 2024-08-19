package main

import (
	"context"
	"fmt"
	"time"
)

func viajar(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Pronto para a viagem!")
	case <-ctx.Done():
		fmt.Println("A viagem foi cancelada:", ctx.Err())
	}
}

func main1() {
	// Cria um contexto com cancelamento
	ctx, cancel := context.WithCancel(context.Background())

	go viajar(ctx)

	// Simula uma decisão após 1 segundo para cancelar a viagem
	time.Sleep(1 * time.Second)
	cancel()

	// Espera um pouco antes de sair
	time.Sleep(2 * time.Second)
}
