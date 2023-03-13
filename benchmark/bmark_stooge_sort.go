package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkStoogeSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go stoogeSort(arreglos[i], wg)
	}
}

func stoogeSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Stooge Sort 10 mil elementos"
	default:
		titulo = "Stooge Sort es mucha demora, omitiendo..."
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.StoogeSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)
}
