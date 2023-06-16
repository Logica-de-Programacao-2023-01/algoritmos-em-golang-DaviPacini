//Faça um algoritmo que imprima os números de 10 a 1 em ordem decrescente.

package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
