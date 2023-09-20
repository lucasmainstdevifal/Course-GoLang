package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("A quantidade de elementos veio com o comando len=%d e capacidade cap=%d %v\n", len(s), cap(s), s)
	// o len(s[:0]) -> Faz com que do inicio do array para direita zere tudo
	fmt.Printf("len = %d cap = %d %v \n", len(s[:0]), cap(s[:0]), s[:0])
	// o : pode ser entendido como em diante, do segundo elemento em diante
	fmt.Printf("len = %d cap = %d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	// Como aumentar a capacidade de um slice
	s = append(s, 110)
	fmt.Printf("%v", s)
}
