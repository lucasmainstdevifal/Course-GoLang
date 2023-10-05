package main

import "fmt"

func SomaInteiro(m map[string]int) int {
	var operation int = 0
	for _, v := range m {
		operation += v
	}

	return operation
}

func somaMaisUm[T int](m T) T {
	var n T
	n += 1

	return n
}

func SomaFloat(m map[string]float64) float64 {
	var operation float64 = 0
	for _, v := range m {
		operation += v
	}

	return operation
}

// Criação de Constraint
type MyNumber int
type Number interface {
	~int | ~float64
}

// Essa função pode somar inteiros ou float64 por causa do Generics T
func Soma[T int | float64](m map[string]T) T {
	var operation T
	for _, value := range m {
		operation += value
	}

	return operation
}

/*func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}*/

// Da erro por que any pode ser qualquer tipo então no if não temos como confirmar se a e b são do mesmo tipo
/*func Compara[T any](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}*/

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	somaUm := map[string]int{"Lucas": 2100, "Ilana": 1000}
	somainteiro := SomaInteiro(somaUm)
	fmt.Printf("O valor mensal é de: %d\n", somainteiro)
	somaDois := map[string]float64{"Lucas": 210.0, "Ilana": 25.52}
	somafloat := SomaFloat(somaDois)
	fmt.Printf("O valor mensal é de: %f\n", somafloat)
	//.....................................................................................
	// Na real é rídiculo criar duas funções iguais só por que os tipos são diferentes né:
	// Então os Generics nos ajudam justamente nessas situações , observe a função Soma
	// Essa função pode somar inteiros ou float64 por causa do Generics -> T :=>
	//.....................................................................................
	// func Soma[T int | float64](m map[string]T) T {
	// 	var operation T
	// 	for _, value := range m {
	// 		operation += value
	// 	}
	//
	// 	return operation
	// }
	//.....................................................................................
	// Se a gente quiser podemos trabalhar com constraints. Que são tipos específicos que a
	// gente cria para ser substituídos
	//.....................................................................................

	// Pelo fato de usar o Generics não importa se colocamos inteiro ou float64 ele irá somar:
	fmt.Println(Soma(somaUm))
	fmt.Println(Soma(somaDois))

	somaTres := map[string]MyNumber{"Lucas": 210, "Ilana": 25}
	fmt.Println(somaTres)

	fmt.Println(Compara(10, 10.0))
}
