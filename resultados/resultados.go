package resultados

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/util"
	"sort"
)

var Resultados []modelos.Resultado

var resultadosBinaryInsertion []modelos.Resultado
var resultadosBitonic []modelos.Resultado
var resultadosBubble []modelos.Resultado
var resultadosBucket []modelos.Resultado
var resultadosGnome []modelos.Resultado
var resultadosHeap []modelos.Resultado
var resultadosInsertion []modelos.Resultado
var resultadosMerge []modelos.Resultado
var resultadosQuick []modelos.Resultado
var resultadosRadix []modelos.Resultado
var resultadosRecursiveInsertion []modelos.Resultado
var resultadosSelection []modelos.Resultado
var resultadosShaker []modelos.Resultado
var resultadosShell []modelos.Resultado
var resultadosStooge []modelos.Resultado
var resultadosStrand []modelos.Resultado

func Consolidar() {
	filtrar()
	ordenar()
	reconstruir()
}

func filtrar() {
	for i := range Resultados {
		switch Resultados[i].Algoritmo {
		case modelos.BINARY_INSERTION_SORT:
			resultadosBinaryInsertion = append(resultadosBinaryInsertion, Resultados[i])
		case modelos.BITONIC_SORT:
			resultadosBitonic = append(resultadosBitonic, Resultados[i])
		case modelos.BUBBLE_SORT:
			resultadosBubble = append(resultadosBubble, Resultados[i])
		case modelos.BUCKET_SORT:
			resultadosBucket = append(resultadosBucket, Resultados[i])
		case modelos.GNOME_SORT:
			resultadosGnome = append(resultadosGnome, Resultados[i])
		case modelos.HEAP_SORT:
			resultadosHeap = append(resultadosHeap, Resultados[i])
		case modelos.INSERTION_SORT:
			resultadosInsertion = append(resultadosInsertion, Resultados[i])
		case modelos.MERGE_SORT:
			resultadosMerge = append(resultadosMerge, Resultados[i])
		case modelos.QUICK_SORT:
			resultadosQuick = append(resultadosQuick, Resultados[i])
		case modelos.RADIX_SORT:
			resultadosRadix = append(resultadosRadix, Resultados[i])
		case modelos.RECURSIVE_INSERTION_SORT:
			resultadosRecursiveInsertion = append(resultadosRecursiveInsertion, Resultados[i])
		case modelos.SELECTION_SORT:
			resultadosSelection = append(resultadosSelection, Resultados[i])
		case modelos.SHAKER_SORT:
			resultadosShaker = append(resultadosShaker, Resultados[i])
		case modelos.SHELL_SORT:
			resultadosShell = append(resultadosShell, Resultados[i])
		case modelos.STOOGE_SORT:
			resultadosStooge = append(resultadosStooge, Resultados[i])
		case modelos.STRAND_SORT:
			resultadosStrand = append(resultadosStrand, Resultados[i])
		}
	}
}

func ordenar() {
	sort.Slice(Resultados, ordenarAscendente(Resultados))

	sort.Slice(resultadosBinaryInsertion, ordenarAscendente(resultadosBinaryInsertion))
	sort.Slice(resultadosBitonic, ordenarAscendente(resultadosBitonic))
	sort.Slice(resultadosBubble, ordenarAscendente(resultadosBubble))
	sort.Slice(resultadosBucket, ordenarAscendente(resultadosBucket))
	sort.Slice(resultadosGnome, ordenarAscendente(resultadosGnome))
	sort.Slice(resultadosHeap, ordenarAscendente(resultadosHeap))
	sort.Slice(resultadosInsertion, ordenarAscendente(resultadosInsertion))
	sort.Slice(resultadosMerge, ordenarAscendente(resultadosMerge))
	sort.Slice(resultadosQuick, ordenarAscendente(resultadosQuick))
	sort.Slice(resultadosRadix, ordenarAscendente(resultadosRadix))
	sort.Slice(resultadosRecursiveInsertion, ordenarAscendente(resultadosRecursiveInsertion))
	sort.Slice(resultadosSelection, ordenarAscendente(resultadosSelection))
	sort.Slice(resultadosShaker, ordenarAscendente(resultadosShaker))
	sort.Slice(resultadosShell, ordenarAscendente(resultadosShell))
	sort.Slice(resultadosStooge, ordenarAscendente(resultadosStooge))
	sort.Slice(resultadosStrand, ordenarAscendente(resultadosStrand))
}

func reconstruir() {
	Resultados = []modelos.Resultado{}
	Resultados = append(Resultados, resultadosBinaryInsertion...)
	Resultados = append(Resultados, resultadosBitonic...)
	Resultados = append(Resultados, resultadosBubble...)
	Resultados = append(Resultados, resultadosBucket...)
	Resultados = append(Resultados, resultadosGnome...)
	Resultados = append(Resultados, resultadosHeap...)
	Resultados = append(Resultados, resultadosInsertion...)
	Resultados = append(Resultados, resultadosMerge...)
	Resultados = append(Resultados, resultadosQuick...)
	Resultados = append(Resultados, resultadosRadix...)
	Resultados = append(Resultados, resultadosRecursiveInsertion...)
	Resultados = append(Resultados, resultadosSelection...)
	Resultados = append(Resultados, resultadosShaker...)
	Resultados = append(Resultados, resultadosShell...)
	Resultados = append(Resultados, resultadosStooge...)
	Resultados = append(Resultados, resultadosStrand...)
}

func ordenarAscendente(arreglo []modelos.Resultado) func(int, int) bool {
	return func(i, j int) bool {
		if arreglo[i].Duracion > arreglo[j].Duracion {
			return arreglo[i].CantidadElementos > arreglo[j].CantidadElementos
		}
		return arreglo[i].CantidadElementos < arreglo[j].CantidadElementos
	}
}

func EscribirResultado() {
	cadena := ""

	for i := range Resultados {
		cadena += Resultados[i].Titulo + " " + Resultados[i].Duracion.String() + "\n"
	}

	util.EscribirArchivo("resultados.txt", []byte(cadena))
}
