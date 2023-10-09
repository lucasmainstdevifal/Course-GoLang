package main

import "fmt"

func main() {
	fmt.Println(sumtwo(3, 4, 5))
}

func sumtwo(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
