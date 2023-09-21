package main

import "fmt"

type Cliente struct {
	nome string
}

// Altera diretamente no local da memória
func (c *Cliente) andou() {
	c.nome = "Lucas P. cabral"
	fmt.Printf("O cliente %v andou\n", c.nome)
}
func main() {
	wesley := Cliente{nome: "Wesley"}
	wesley.andou()
	fmt.Printf("O valor da struct é %v", wesley.nome)
}
