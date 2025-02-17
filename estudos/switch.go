package main

import "fmt"

func Switch() {

var dia string = "segunda"

switch dia {
case "segunda":
	fmt.Println("Hoje é Segunda")
case "terça", "quarta", "quinta", "sexta":
	fmt.Println("Hoje é Terça, Quarta, Quinta ou Sexta")
default:
	fmt.Println("Outro dia da semana.")
}


}