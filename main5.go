//5. Faça um algoritmo que leia o valor do dólar em reais e um valor em dólares, e converta
//esse valor para reais.

package main

import "fmt"

func main() {
	var dolar float32
	var real float32
	var nome string
	fmt.Println("Ola, qual seu nome? ") //pergunta nominal
	fmt.Scan(&nome)
	fmt.Println("Ola Sr(a).", nome, "qual seria o valor em Dolar para realizar o cambio?")
	fmt.Print("$")
	fmt.Scan(&dolar)
	real = dolar * 5.23
	fmt.Print("O valor retornado sera de R$ ", real)

}
