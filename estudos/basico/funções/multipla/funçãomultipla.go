package main

import "fmt"

func calculo(){
	var a int = 10
	var b int = 5
	var soma int = a + b
	var multiplicacao int = a * b
	fmt.Println("Resultado: ", soma, multiplicacao)
}

func dividir(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Não é possível dividir por zero.")
	}
	return a / b, nil
}

func subtrair() {
	var a int = 10
	var b int = 5
	var resultado int = a - b
	fmt.Println("Resultado: ", resultado)
}


func main() {
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Resultado da divisão: ", resultado)
	}
}