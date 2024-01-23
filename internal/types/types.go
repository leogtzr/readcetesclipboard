package types

import "fmt"

const (
	goal = 1_800_000
)

type Monto struct {
	Invertido  float64
	PlusMinus  float64
	Disponible float64
	Valuado    float64
}

func (m Monto) String() string {
	return fmt.Sprintf("Monto Invertido=%.2f\tDisponible=%.2f\tValuado=%.2f", m.Invertido, m.Disponible, m.Valuado)
}
