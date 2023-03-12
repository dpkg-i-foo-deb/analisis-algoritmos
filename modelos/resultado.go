package modelos

import "time"

type Resultado struct {
	Titulo            string        `json:"nombre"`
	CantidadElementos int           `json:"cantidadElementos"`
	Duracion          time.Duration `json:"duracion"`
}
