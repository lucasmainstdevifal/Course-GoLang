package main

import "fmt"

func soma(a, b int) int {

	return a + b
}

func somaPonteiro(a, b *int) int {
	*a = 50
	*b = 55
	return *a + *b
}

func main() {
	valor1 := 10
	valor2 := 20
	fmt.Println(soma(valor1, valor2), "\n")

	fmt.Println(valor1, "\n")
	fmt.Println(valor2, "\n")
	//======================================
	fmt.Println(somaPonteiro(&valor1, &valor2))
	fmt.Println(valor1)
	fmt.Println(valor2)

}
