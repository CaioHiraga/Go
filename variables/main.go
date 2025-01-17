package main

// programa escopo 1 (inicial/maior escopo)
import "fmt"

const i float64 = 5.5

// main - escopo 2 (menor / mais limitado)
func main() {
	x := "Hello World" //criação/declaração e atribuição de x -> :=
	fmt.Println(x)     //referencia a x

	var y string //criação/declaração de y
	y = "y: Hello World"
	fmt.Println(y) //referencia a y

	var z, a, b int = 1, 2, 3 //criacao/declaração de n variáveis
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	c, d := 3, 5
	fmt.Println(c)
	fmt.Println(d)

	// estaticamente tipado + fortemente tipado

	fmt.Println(i)
}
