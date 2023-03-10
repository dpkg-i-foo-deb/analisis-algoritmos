package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkStoogeSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		stoogeSort(arreglos[i])
	}
}

func stoogeSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Stooge Sort 10 mil elementos"
	case 100000:
		titulo = "Stooge Sort 100 mil elementos"
	case 1000000:
		titulo = "Stooge Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.StoogeSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)
}
