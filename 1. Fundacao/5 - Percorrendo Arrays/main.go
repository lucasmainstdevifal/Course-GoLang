package main

import "fmt"

func main() {
	// Aqui definimos um array de 3 posições e esse tamanho é fixo
	var meuArray [3]int

	// Aqui estamos atribuindo valores a cada posição do array
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 15

	fmt.Println(len(meuArray) - 1)         // Retornara o último índice do array
	fmt.Println(meuArray[len(meuArray)-1]) // Retornará o valor que está no elemento de último índice do array

	// Para Percorrer um array usamos o for do Go com a seguinte estrutura:
	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d.\n", i, v)
	}
}
