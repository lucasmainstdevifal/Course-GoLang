package main

import "fmt"

func main() {
	total := func() int {
		return sumtwo(1, 2, 3, 4, 5, 6, 7) * 2
	}()

	fmt.Println(total)
}

func sumtwo(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
