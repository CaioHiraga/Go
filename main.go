package main

import "fmt"

func main() {
	fmt.Println("Hello, Go Essentials!") //String

	fmt.Println(2) //Inteiros
	fmt.Printf("Hello, I'm %v and I'm an integer\n", 2)
	fmt.Print("Hello, I'm %d and I'm an integer\n", 2)

	fmt.Println(2 + 5) // Expressões

	//Booleanos
	fmt.Println(true)
	fmt.Println(false)

}

// Pasta == folders podem ter múltiplos arquivos
// Pasta pode ter multiplas funções dentro dela (espalhadas em múltiplos arquivos)
// Toda pasta é um pacote
// Funções pertencem a um pacote
