//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.

package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		num3 int
		nome string
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Informe o primeiro número, ", nome)
	fmt.Scan(&num1)
	fmt.Println("Informe o segundo número")
	fmt.Scan(&num2)
	fmt.Println("Informe o terceiro e último número")
	fmt.Scan(&num3)
	soma := num1 + num2 + num3
	fmt.Print("A soma dos números: ", num1, ", ", num2, " e ", num3, " é ", soma)

}
