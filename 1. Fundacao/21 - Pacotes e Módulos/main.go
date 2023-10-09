package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println("Resultado: %v", s)
	fmt.Println(carro)
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
}
