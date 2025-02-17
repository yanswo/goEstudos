package main

import "fmt"

func main() {
	// Arrays são estruturas de dados que armazenam uma coleção de elementos do mesmo tipo.
	var numeros [5]int = [5]int{2, 4, 6, 8, 10}

	// Acessando elementos do array
	fmt.Println(numeros[0])
	fmt.Println(numeros[4])

	// Modificando
	numeros[2] = 66
	fmt.Println("Resultado dos numeros com modificação: ", numeros)


	// Percorrendo o array	
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}
}