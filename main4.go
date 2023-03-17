package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var num3 float64
	var num4 float64
	var nome string

	fmt.Println("Ola, qual seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("Ola", nome, ", por favor informe 4 numeros.")
	fmt.Println(nome, "informe o primeiro numero")
	fmt.Scan(&num1)
	fmt.Println("Informe o segundo numero")
	fmt.Scan(&num2)
	fmt.Println("Informe o terceiro numero")
	fmt.Scan(&num3)
	fmt.Println("Informe o ultimo numero")
	fmt.Scan(&num4)
	fmt.Print("A media aritimetica dos numeros: ", num1, ", ", num2, ", ", num3, " e ", num4, " é ", (num1+num2+num3+num4)/4)

}
