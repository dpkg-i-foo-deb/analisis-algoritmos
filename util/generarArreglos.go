package util

import "math/rand"

func GenerarArreglo(cantidad int) []int {

	resultado := make([]int, 0)

	for i := 0; i < cantidad; i++ {
		valor := rand.Intn(9000) + 1000
		resultado = append(resultado, valor)
	}

	return resultado
}
