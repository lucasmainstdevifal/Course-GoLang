package main

import (
	"errors"
	"fmt"
)

func main() {
	h := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("Temos a quantidade de %d e a capacidade de %d o próprio array %v\n", len(h), cap(h), h)

	s, e := sum(30, 40)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(s)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50.\n")
	}
	return a + b, nil
}
