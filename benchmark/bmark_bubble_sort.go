package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkBubbleSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		bubbleSort(arreglos[i])
	}
}

func bubbleSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Bubble Sort 1 millón de elementos"
	case 10000000:
		titulo = "Bubble Sort 10 millones de elementos"
	case 100000000:
		titulo = "Bubble Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BubbleSort(&arreglo.Arr)
}
