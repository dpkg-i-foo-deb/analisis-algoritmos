package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkSelectionSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go selectionSort(arreglos[i], wg)
	}
}

func selectionSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.SELECTION_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.SelectionSort(&arreglo.Arr)
}
