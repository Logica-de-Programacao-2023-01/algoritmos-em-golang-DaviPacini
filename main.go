// Faça um algoritmo que leia o nome, a idade e o peso de uma pessoa e imprima esses
// valores na tela.
package main

import (
	"fmt"
)

func main() {
	var nome string
	var peso float64
	var idade uint
	fmt.Print("Qual seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("Ok,", nome, "quantos anos voce tem?")
	fmt.Scan(&idade)
	if idade >= 18 {
		fmt.Println("Qual seu peso? ")
		fmt.Scan(&peso)
		fmt.Println("Suas informacoes: ")
		fmt.Println("Nome:", nome)
		fmt.Println("Idade:", idade, "Anos")
		fmt.Println("Peso:", peso, "Kg")
	}
	if idade < 18 {
		fmt.Println("ACESSO NEGADO")
	}
}
