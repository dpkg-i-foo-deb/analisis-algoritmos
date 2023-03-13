package modelos

import "time"

type AlrogitmoOrdenamiento string

const (
	BINARY_INSERTION_SORT    AlrogitmoOrdenamiento = "Binary Insertion Sort"
	BITONIC_SORT             AlrogitmoOrdenamiento = "Bitonic Sort"
	BUBBLE_SORT              AlrogitmoOrdenamiento = "Bubble Sort"
	BUCKET_SORT              AlrogitmoOrdenamiento = "Bucket Sort"
	GNOME_SORT               AlrogitmoOrdenamiento = "Gnome Sort"
	HEAP_SORT                AlrogitmoOrdenamiento = "Heap Sort"
	INSERTION_SORT           AlrogitmoOrdenamiento = "Insertion Sort"
	MERGE_SORT               AlrogitmoOrdenamiento = "Merge Sort"
	QUICK_SORT               AlrogitmoOrdenamiento = "Quick Sort"
	RADIX_SORT               AlrogitmoOrdenamiento = "Radix Sort"
	RECURSIVE_INSERTION_SORT AlrogitmoOrdenamiento = "Recursive Insertion Sort"
	SELECTION_SORT           AlrogitmoOrdenamiento = "Selection Sort"
	SHAKER_SORT              AlrogitmoOrdenamiento = "Shaker Sort"
	SHELL_SORT               AlrogitmoOrdenamiento = "Shell Sort"
	STOOGE_SORT              AlrogitmoOrdenamiento = "Stooge Sort"
	STRAND_SORT              AlrogitmoOrdenamiento = "Strand Sort"
)

type Resultado struct {
	Titulo            string `json:"nombre"`
	Algoritmo         AlrogitmoOrdenamiento
	CantidadElementos int           `json:"cantidadElementos"`
	Duracion          time.Duration `json:"duracion"`
}
