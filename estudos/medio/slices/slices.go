package main

import "fmt"

func main() {
	// Declaração de um slice
	var frutas = []string{"banana", "maçã", "uva", "morango", "abacaxi"}
	fmt.Println(frutas)

	// Adicionando um elemento ao slice
	frutas = append(frutas, "laranja", "melancia")
	fmt.Println(frutas)
}