package util

import (
	"fmt"
	"strconv"
)

func ImprimirMatriz(matriz [][]int) {
	cadena := ""

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			cadena += strconv.FormatInt(int64(matriz[i][j]), 10)
			cadena += "   "
		}
		cadena += "\n"
	}

	fmt.Println(cadena)
}
