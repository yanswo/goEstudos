package main

import "fmt"

func main() {
	// Declaração de um map
	pessoas:= map[string]int{
		"João": 20,
		"Maria": 25,
		"José": 30,
	}
	// Acessando elementos do map
	fmt.Println(pessoas["João"]) 
	pessoas["João"] = 21 // Modificando
	fmt.Println(pessoas["João"])

	valor, existe := pessoas["Maria"]
	if existe {
		fmt.Println("Maria existe e tem", valor, "anos")
	} else {
		fmt.Println("Maria não existe")
	}


	pessoas1 := map[string]int{ 
		"Yan": 21,
		"Harrison": 25,
		"Swoot": 31,
	}

	fmt.Println(pessoas1["Yan"])
	pessoas1["Yan"] = 22
	fmt.Println(pessoas1["Yan"])

	valor1, existe1 := pessoas1["Harrison"] 
		if existe1 {
			fmt.Println("Harrison existe e tem", valor1, "anos")
		} else {
			fmt.Println("Harrison não existe")
		}
	

}