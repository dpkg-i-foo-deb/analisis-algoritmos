package pruebas

import (
	"analisis-algoritmos/ordenamiento"
	"reflect"
	"testing"
)

func TestStoogeSort(t *testing.T) {
	esperado := []int{0, 1, 2, 3, 4}
	resultado := []int{4, 2, 1, 3, 0}

	ordenamiento.StoogeSort(&resultado, 0, len(resultado)-1)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Fail()
	}
}
