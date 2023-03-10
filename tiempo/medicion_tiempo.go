package tiempo

import (
	"fmt"
	"time"
)

func MedirTiempo(titulo string) func() time.Duration {
	inicio := time.Now()
	return func() time.Duration {
		fmt.Printf("%s Tiempo %v\n", titulo, time.Since(inicio))

		return time.Since(inicio)
	}
}
