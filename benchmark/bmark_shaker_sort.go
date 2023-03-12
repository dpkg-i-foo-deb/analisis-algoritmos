package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkShakerSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go shakerSort(arreglos[i])
	}
}

func shakerSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Shaker Sort 10 mil elementos"
	case 100000:
		titulo = "Shaker Sort 100 mil elementos"
	case 1000000:
		titulo = "Shaker Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.ShakerSort(&arreglo.Arr)
}
