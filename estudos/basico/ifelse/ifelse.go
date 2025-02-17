package main

import "fmt"

func main() {
	
	var idade int = 21

	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else if idade >= 13 {
		fmt.Println("Adolescente")
	} else {
		fmt.Println("Criança")
	}
	
	
}