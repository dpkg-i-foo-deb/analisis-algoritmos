package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBubbleSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go bubbleSort(arreglos[i], wg)
	}
}

func bubbleSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.BUBBLE_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.BubbleSort(&arreglo.Arr)
}
