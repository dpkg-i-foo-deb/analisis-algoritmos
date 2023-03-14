package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBinaryInsertionSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go binaryInsertionSort((arreglos)[i], wg)
	}
}

func binaryInsertionSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.BINARY_INSERTION_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.BinaryInsertionSort(&arreglo.Arr)
}
