package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"readcetesclipboard/internal/utils"
	"strconv"
)

var (
	financeInvestmentGoal = os.Getenv("CETES_INVESTMENT_GOAL")
)

func init() {
	if financeInvestmentGoal == "" {
		log.Fatal("CETES_INVESTMENT_GOAL environment variable not set")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var texto string
	for scanner.Scan() {
		texto += scanner.Text() + "\n"
	}

	if scanner.Err() != nil {
		fmt.Println("Error al leer la entrada estándar:", scanner.Err())
		return
	}

	monto, err := utils.FromCETESInputClipboardText(texto)

	if err != nil {
		log.Fatal(err)
	}

	goal, err := strconv.ParseFloat(financeInvestmentGoal, 64)
	if err != nil {
		log.Fatal(fmt.Errorf("error: expecting a numeric value in CETES_INVESTMENT_GOAL environment value"))
	}

	fmt.Printf("%s\n\t\t\t\t\t\t\t%.2f%%\n\n", monto, (monto.Valuado*100.0)/goal)
}
