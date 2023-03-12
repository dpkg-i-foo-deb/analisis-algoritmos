package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkBubbleSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go bubbleSort(arreglos[i])
	}
}

func bubbleSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Bubble Sort 10 mil elementos"
	case 100000:
		titulo = "Bubble Sort 100 mil elementos"
	case 1000000:
		titulo = "Bubble Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BubbleSort(&arreglo.Arr)
}
