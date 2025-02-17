package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func main() {
	var resultado int = somar(10, 20)
	fmt.Println("Resultado da soma", resultado)
}