package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBitonicSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go bitonicSort(arreglos[i], wg)
	}
}

func bitonicSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Bitonic Sort 10 mil elementos"
	case 100000:
		titulo = "Bitonic Sort 100 mil elementos"
	case 1000000:
		titulo = "Bitonic Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.BitonicSort(&arreglo.Arr, 0, len(arreglo.Arr), true)
}
