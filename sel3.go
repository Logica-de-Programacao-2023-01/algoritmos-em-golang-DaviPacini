//Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.

package main

import "fmt"

func main() {
	//variável
	var num int
	//estrutura
	fmt.Println("Olá, informe um número")
	fmt.Scan(&num)
	//estruturas if
	if (num % 2) != 0 {
		fmt.Println("O número ", num, " é impar")
	} else if num == 0 {
		fmt.Print("O número ", num, " é par e nulo")
	} else {
		fmt.Println("O número ", num, " é par")
	}
}
