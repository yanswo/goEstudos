package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anônima executada.")
	}() // Chama a função anônima
}