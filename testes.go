package main

import "fmt"

func ma() {
	
	var nome string = "Yan"
	var idade int = 21
	var peso float64 = 73.3
	var vivo bool = true
	
	fmt.Println("Nome: ", nome, "Idade: ", idade, "Peso: ", peso, "Vivo: ", vivo) 

var idadeUsuario int = 13

if idadeUsuario < 12 {
	println("Criança")
} else if idadeUsuario <= 17 {
	println("Adolescente")
} else if idadeUsuario <= 64 {
	println("Adulto")
} else {
	println("Idoso")
}

frutas := []string{"banana", "maçã", "uva", "laranja"}

for _, fruta := range frutas {
	fmt.Println(fruta)
}

}