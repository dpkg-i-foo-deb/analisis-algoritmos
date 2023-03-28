package util

import "fmt"

func ImprimirArregloNumeros(a []int) {
	fmt.Println()
	for i := range a {
		fmt.Print(a[i])
	}
	fmt.Println()
}
