//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento
//de 10% se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

package main

import "fmt"

func main() {
	var (
		salario      float64
		novo_salario float64
	)
	fmt.Println("Informe o salário")
	fmt.Scan(&salario)
	if salario < 1000 {
		novo_salario = salario + (salario * 10 / 100)
	} else {
		novo_salario = salario + (salario * 5 / 100)
	}
	fmt.Printf("O novo salário é de R$ %.2f", novo_salario)
}
