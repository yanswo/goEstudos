package main

import "fmt"

func For() {

// começa do 0, menor que 10 e aumenta em 1 o valor.	
for i := 0; i < 10; i++ {
	fmt.Println("Número: ", i)
}

 // FOR com While
i := 0
for i < 5 {
	fmt.Println("Número: ", i)
	i++ 
}

// Loop Infinito só se omitir tudo do FOR
for {
	fmt.Println("Loop Infinito!")
	break
}

numeros := []int{1, 2, 3, 4, 5}

for index, valor:= range numeros {
	fmt.Printf("Índice: %d, Valor: %d", index, valor)
}

}