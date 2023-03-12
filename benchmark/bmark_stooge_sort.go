package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"log"
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
	case 100000:
		log.Println("Stooge Sort tomaría 21 años con 100 mil elementos")
		return
	case 1000000:
		log.Println("Stooge Sort tomaría 51 millones de años con 1 millón de elementos")
		return
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.StoogeSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)
}
