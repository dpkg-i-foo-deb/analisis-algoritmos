package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkBinaryInsertionSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		binaryInsertionSort((arreglos)[i])
	}
}

func binaryInsertionSort(arreglo *modelos.Arreglo) {

	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Binary Insertion Sort 10 mil elementos"
	case 100000:
		titulo = "Binary Insertion Sort 100 mil elementos"
	case 1000000:
		titulo = "Binary Insertion 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BinaryInsertionSort(&arreglo.Arr)
}
