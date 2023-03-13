package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkRadixSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go radixSort(arreglos[i], wg)
	}
}

func radixSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Radix Sort 10 mil elementos"
	case 50000:
		titulo = "Radix Sort 50 mil elementos"
	case 100000:
		titulo = "Radix Sort 100 mil elementos"
	case 500000:
		titulo = "Radix Sort 500 mil elementos"
	case 1000000:
		titulo = "Radix Sort 1 millón de elementos"
	case 10000000:
		titulo = "Radix Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.RadixSort(&arreglo.Arr)
}
