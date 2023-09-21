package main

import "fmt"

func main() {
	var minhavar interface{} = "Wesley Willians"

	fmt.Println(minhavar.(string))
}
