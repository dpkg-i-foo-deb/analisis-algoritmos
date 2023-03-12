package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkRadixSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go radixSort(arreglos[i])
	}
}

func radixSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Radix Sort 10 mil elementos"
	case 100000:
		titulo = "Radix Sort 100 mil elementos"
	case 1000000:
		titulo = "Radix Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.RadixSort(&arreglo.Arr)
}
