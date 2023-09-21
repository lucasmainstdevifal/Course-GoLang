package main

func main() {

	// Memória -> Endereço -> Valor
	// Variável -> Endereço -> Valor
	a := 10
	//fmt.Println(a)
	//fmt.Println(&a)

	var ponteiro *int = &a
	*ponteiro = 20
	println(a)
	println(*ponteiro)

	b := &a
	*b = 30
	println(*b)
	println(a)
}
