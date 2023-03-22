//Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.

package main

import "fmt"

func main() {
	var (
		nome string
		kg   float64
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println(nome, "informe seu peso em quilos para saber a converção em libras")
	fmt.Print("Kg:")
	fmt.Scan(&kg)
	lbs := kg * 2.2
	fmt.Print("Seu peso em Libras é de ", lbs, "lbs")
}
