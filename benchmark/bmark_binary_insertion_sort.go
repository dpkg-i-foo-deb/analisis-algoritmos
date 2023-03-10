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
	case 1000000:
		titulo = "Binary Insertion Sort 1 millón de elementos"
	case 10000000:
		titulo = "Binary Insertion Sort 10 millones de elementos"
	case 100000000:
		titulo = "Binary Insertion Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BinaryInsertionSort(&arreglo.Arr)
}
