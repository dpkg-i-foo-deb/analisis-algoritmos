package tiempo

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/resultados"
	"fmt"
	"time"
)

func MedirTiempo(titulo string, algoritmo modelos.AlrogitmoOrdenamiento, cantidadElementos int) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		fmt.Printf("%s Tiempo %v\n", titulo, time.Since(inicio))

		resultado := modelos.Resultado{
			Titulo:            titulo,
			Algoritmo:         algoritmo,
			CantidadElementos: cantidadElementos,
			Duracion:          time.Since(inicio),
		}

		resultados.Resultados = append(resultados.Resultados, resultado)

		return time.Since(inicio)
	}
}
