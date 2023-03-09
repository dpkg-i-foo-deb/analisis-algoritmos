package util

import "math/rand"

func GenerarArreglo(digitos int, cantidad int) []int {

	minimo := int64(10 ^ (digitos - 1))
	maximo := int64((10 ^ digitos) - 1)

	resultado := make([]int, 0)

	for i := 0; i < cantidad; i++ {
		valor := int(minimo + rand.Int63n(maximo-minimo+1))
		resultado = append(resultado, valor)
	}

	return resultado
}
