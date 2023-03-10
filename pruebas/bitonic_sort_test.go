package pruebas

import (
	"analisis-algoritmos/ordenamiento"
	"fmt"
	"reflect"
	"testing"
)

func TestBitonicSort(t *testing.T) {
	esperado := []int{0, 1, 2, 3, 4}
	resultado := []int{4, 2, 1, 3, 0}

	ordenamiento.BitonicSort(&resultado, 0, len(resultado), true)

	fmt.Print(resultado)

	if !reflect.DeepEqual(resultado, esperado) {
		t.Fail()
	}
}
