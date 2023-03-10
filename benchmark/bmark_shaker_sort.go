package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkShakerSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		shakerSort(arreglos[i])
	}
}

func shakerSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Shaker Sort 1 millón de elementos"
	case 10000000:
		titulo = "Shaker Sort 10 millones de elementos"
	case 100000000:
		titulo = "Shaker Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.ShakerSort(&arreglo.Arr)
}
