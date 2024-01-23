package utils

import (
	"fmt"
	"readcetesclipboard/internal/types"
	"strconv"
	"strings"
)

const (
	mandatoryInputStringStartPointText = "TOTAL"
)

func hasStartingInputTextPoint(input string) (bool, int) {
	if len(input) == 0 {
		return false, -1
	}

	lines := strings.Split(input, "\n")

	has := false
	idx := -1

	for i, line := range lines {
		if line == mandatoryInputStringStartPointText {
			has = true
			idx = i

			break
		}
	}

	return has, idx
}

func validateMontosTextLines(montosLine, montoValuado string) error {
	if len(montosLine) == 0 || len(strings.TrimSpace(montosLine)) == 0 {
		return fmt.Errorf("montos line is empty")
	}

	montos := strings.Fields(strings.TrimSpace(montosLine))

	if len(montos) != 3 {
		return fmt.Errorf("not enough montos values: %s", montosLine)
	}

	for _, monto := range montos {
		monto = strings.ReplaceAll(monto, ",", "")
		_, err := strconv.ParseFloat(monto, 64)

		if err != nil {
			return err
		}
	}

	montoValuado = strings.ReplaceAll(montoValuado, ",", "")

	if _, err := strconv.ParseFloat(montoValuado, 64); err != nil {
		return err
	}

	return nil
}

func setFromMontosLine(monto *types.Monto, montosLine, montoValuadoLine string) error {
	montos := strings.Fields(strings.TrimSpace(montosLine))

	montoInvertido, err := strconv.ParseFloat(montos[0], 64)
	if err != nil {
		return err
	}
	monto.Invertido = montoInvertido

	plusMinus, err := strconv.ParseFloat(montos[1], 64)
	if err != nil {
		return err
	}
	monto.PlusMinus = plusMinus

	montoDisponible, err := strconv.ParseFloat(montos[2], 64)
	if err != nil {
		return err
	}
	monto.Disponible = montoDisponible

	montoValuado, err := strconv.ParseFloat(montoValuadoLine, 64)
	if err != nil {
		return err
	}

	monto.Valuado = montoValuado

	return nil
}

func FromCETESInputClipboardText(input string) (types.Monto, error) {
	has, startingPointIdx := hasStartingInputTextPoint(input)

	if !has {
		return types.Monto{}, fmt.Errorf("input text does not contains %s", mandatoryInputStringStartPointText)
	}

	lines := strings.Split(input, "\n")

	if (startingPointIdx + 3) >= len(lines) {
		return types.Monto{}, fmt.Errorf("not enough lines to parse, %s+2 lines", mandatoryInputStringStartPointText)
	}

	montosLine := lines[startingPointIdx+2]
	montoValuadoLine := lines[startingPointIdx+3]

	if err := validateMontosTextLines(montosLine, montoValuadoLine); err != nil {
		fmt.Println("here...")
		return types.Monto{}, err
	}

	monto := types.Monto{}

	montosLine = strings.ReplaceAll(montosLine, ",", "")
	montoValuadoLine = strings.ReplaceAll(montoValuadoLine, ",", "")

	err := setFromMontosLine(&monto, montosLine, montoValuadoLine)

	if err != nil {
		return types.Monto{}, err
	}

	return monto, nil
}
