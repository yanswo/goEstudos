package main

import (
	"fmt"
	"time"
)

// Channels são canais que permitem a comunicação entre goroutines.
func tarefa(nome string, ch chan string) {
// Enviando um valor para o canal
	fmt.Println(nome + " iniciada.")
	time.Sleep(2 * time.Second) 
	ch <- nome + " finalizada." 
}

func main() {
	// Criando um canal
	ch := make(chan string)


	go tarefa("Tarefa 1", ch)
	
// Recebendo o valor do canal
	fmt.Println(<-ch)

	
	go tarefa("Tarefa 2", ch)
	

	fmt.Println(<-ch)

	fmt.Println("Fim do programa principal")
}
