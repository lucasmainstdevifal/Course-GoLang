package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome   string
	Idade  int
	Ativo  bool
	Adress Endereco
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func Desativacao(pesssoa Pessoa) {
	pesssoa.Desativar()
}

func main() {
	Lucas := Cliente{
		Nome:  "Lucas",
		Idade: 27,
		Ativo: true,
	}

	Lucas.Ativo = false
	Lucas.Adress.Cidade = "Palmeira dos Índios"
	Lucas.Adress.Cidade = "Maceio -Al"

	fmt.Printf("Nome: %s, Idade %d, Ativo %t", Lucas.Nome, Lucas.Idade, Lucas.Ativo)
	Lucas.Desativar()
	fmt.Println("O cliente foi desativado \n", Lucas.Ativo)

	Desativacao(Lucas)
}
