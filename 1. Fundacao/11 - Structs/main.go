package main

import "fmt"

// Isso aqui é uma Struct
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	Lucas := Cliente{
		Nome:  "Lucas",
		Idade: 27,
		Ativo: true,
	}

	Lucas.Ativo = false

	fmt.Printf("Nome: %s, Idade %d, Ativo %t", Lucas.Nome, Lucas.Idade, Lucas.Ativo)
	fmt.Println(Lucas.Ativo)
}
