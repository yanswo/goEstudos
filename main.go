package main

import "fmt"



func main() {
	anonima := func() {
		fmt.Println("Olá eu sou uma função anônima!.")
	}

	anonima()
}