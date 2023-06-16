//Faça um algoritmo que imprima os números de 1 a 10 em ordem crescente.

package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
