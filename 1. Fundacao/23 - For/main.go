package main

import "fmt"

func main() {
	// No go temos apenas o FOR
	for i := 0; i < 10; i++ {
		fmt.Println("%d", i)
	}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}
}
