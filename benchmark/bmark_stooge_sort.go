package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"log"
)

func BmarkStoogeSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		stoogeSort(arreglos[i])
	}
}

func stoogeSort(arreglo modelos.Arreglo) {
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

	ordenamiento.StoogeSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)
}
