package main

import (
	"errors"
	"fmt"
)

func main() {
	//sub(2, 3)

	valor, err := sumfour(49, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

func sum(a int, b int) int {
	return a + b
}

func sumtwo(a, b int) int {
	return a + b
}

func sumthree(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

func sumfour(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}

func sub(a int, b int) {
	sub := a - b
	fmt.Printf("%d", sub)
}
