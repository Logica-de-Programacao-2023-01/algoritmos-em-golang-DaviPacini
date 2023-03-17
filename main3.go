//Faça um algoritmo que calcule o volume de uma caixa com base, altura e profundidade
//informados pelo usuário.

package main

import "fmt"

func main() {
	var base float64
	var altura float64
	var profundidade float64
	var nome string
	fmt.Println("Qual seu nome? ")
	fmt.Scan(&nome)
	fmt.Print("Ola ", nome, " informe o valor da base, em CM: ")
	fmt.Scan(&base)
	if base <= 0 {
		fmt.Println("Valor invalido, tente novamente")
		fmt.Print("Base em CM: ")
		fmt.Scan(&base)
	}
	fmt.Print("Informe a altura, em CM: ")
	fmt.Scan(&altura)
	if altura <= 0 {
		fmt.Println("Valor invalido, tente novamente")
		fmt.Print("altura em CM: ")
		fmt.Scan(&altura)
	}
	fmt.Print("Profundidade, em CM: ")
	fmt.Scan(&profundidade)
	if profundidade <= 0 {
		fmt.Println("Valor invalido, tente novamente")
		fmt.Print("profundidade em CM: ")
		fmt.Scan(&profundidade)
	}
	volume := base * altura * profundidade
	fmt.Println("O volume da caixa é de", volume, "CM^3")
}
