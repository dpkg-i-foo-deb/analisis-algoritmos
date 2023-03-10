package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkRadixSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		radixSort(arreglos[i])
	}
}

func radixSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Radix Sort 1 millón de elementos"
	case 10000000:
		titulo = "Radix Sort 10 millones de elementos"
	case 100000000:
		titulo = "Radix Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.RadixSort(&arreglo.Arr)
}
