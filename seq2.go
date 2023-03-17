package main

import "fmt"

func main() {
	var (
		nome string
		num  float64
	)
	fmt.Println("Olá, qual seu nome? ")
	fmt.Scan(&nome)
	fmt.Println(nome, " informe um número para ser calculado seu dobro, triplo e quadruplo")
	fmt.Scan(&num)
	dobro := num * 2
	triplo := num * 3
	quadruplo := num * 4
	fmt.Println("O dobro de ", num, " é ", dobro)
	fmt.Println("O triplo de ", num, " é ", triplo)
	fmt.Println("O quadruplo de ", num, " é ", quadruplo)
}
