package main

import "fmt"

// Função que recebe uma função como argumento
func executarFuncao(f func()) {
	f()
}


// Função principal
func main() {
	minhaFuncao := func() {
		fmt.Println("Função passada como argumento!.")
	}
	executarFuncao(minhaFuncao)
}