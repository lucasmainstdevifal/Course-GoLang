package main

import "fmt"

func main() {
	var minhavar interface{} = "Wesley Willians"

	fmt.Println(minhavar.(string))
	res, ok := minhavar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
}
