package tiempo

import (
	"fmt"
	"time"
)

func MedirTiempo(titulo string) func() {
	inicio := time.Now()
	return func() {
		fmt.Printf("%s Tiempo %v\n", titulo, time.Since(inicio))
	}
}
