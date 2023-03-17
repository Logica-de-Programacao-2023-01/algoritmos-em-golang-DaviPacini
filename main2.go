// Faça um algoritmo que calcule a área de um retângulo com base e altura informados
// pelo usuário.
package main

import (
	"fmt"
)

func main() {
	var base float64
	var altura float64
	fmt.Print("Informe o tamanho da base em CM: ")
	fmt.Scan(&base)
	if base <= 0 {
		fmt.Println("Tente novamente com valores validos")
	}
	fmt.Println("Informe a altura: ")
	fmt.Scan(&altura)
	if altura <= 0 {
		fmt.Println("Tente novamente com valores validos")
	}
	area := base * altura
	fmt.Println("a area do retangulo de base", base, "CM", "e altura", altura, "CM", "é", area, "CM")

}
