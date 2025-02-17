package main

import (
	"fmt"
	"time"
)

// Goroutines são funções que são executadas de forma concorrente.

func saudacao(nome1 string,) {
	fmt.Println("Olá " + nome1)
	
}

// A função main é a entrada de um programa em Go.

func tarefa(nome string, duracao time.Duration) {
	fmt.Println("Iniciando tarefa", nome, 1 * time.Second)
	time.Sleep(duracao)
	fmt.Println("Tarefa", nome, "finalizada!")
}

func main() {
	go saudacao("Yan")

	go tarefa("Tarefa 1", 2 * time.Second)
	go tarefa("Tarefa 2", 3 * time.Second)
	go tarefa("Tarefa 3", 4 * time.Second)

	time.Sleep(6 * time.Second)
}