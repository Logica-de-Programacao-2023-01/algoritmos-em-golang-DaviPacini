//Faça um algoritmo que imprima os múltiplos de 3 de 0 a 30.

package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 0
	for i <= 30 {
		fmt.Println(i)
		i += 3
	}
}
