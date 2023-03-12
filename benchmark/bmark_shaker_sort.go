package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkShakerSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go shakerSort(arreglos[i], wg)
	}
}

func shakerSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Shaker Sort 10 mil elementos"
	case 100000:
		titulo = "Shaker Sort 100 mil elementos"
	case 1000000:
		titulo = "Shaker Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.ShakerSort(&arreglo.Arr)
}
