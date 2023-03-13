package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkStrandSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go strandSort(arreglos[i], wg)
	}
}

func strandSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Strand Sort 10 mil elementos"
	case 50000:
		titulo = "Strand Sort 50 mil elementos"
	case 100000:
		titulo = "Strand Sort 100 mil elementos"
	case 500000:
		titulo = "Strand Sort 500 mil elementos"
	case 1000000:
		titulo = "Strand Sort 1 millón de elementos"
	case 10000000:
		titulo = "Strand Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo, modelos.STRAND_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.StrandSort(&arreglo.Arr)
}
