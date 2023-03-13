package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBubbleSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		go bubbleSort(arreglos[i], wg)
	}
}

func bubbleSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Bubble Sort 10 mil elementos"
	case 50000:
		titulo = "Bubble Sort 50 mil elementos"
	case 100000:
		titulo = "Bubble Sort 100 mil elementos"
	case 500000:
		titulo = "Bubble Sort 500 mil elementos"
	case 1000000:
		titulo = "Bubble Sort 1 millón de elementos"
	case 10000000:
		titulo = "Bubble Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.BubbleSort(&arreglo.Arr)
}
