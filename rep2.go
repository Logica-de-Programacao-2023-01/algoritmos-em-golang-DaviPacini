//Faça um algoritmo que imprima os números pares de 0 a 20.

package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 0
	for i <= 20 {
		fmt.Println(i)
		i += 2
	}
}
